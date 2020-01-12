package handler

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"shortlyst/pkg/model"
	"shortlyst/pkg/service"
	"strconv"
	"strings"
)

type handler struct {
	item  service.ItemService
	saldo service.SaldoService
}

// NewHandler ds
func NewHandler(item service.ItemService, saldo service.SaldoService) handler {
	return handler{
		item:  item,
		saldo: saldo,
	}
}

func (hndl *handler) Start() bool {
	fmt.Print("Input 1 untuk import dan start dan 2 untuk start saja : ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	menu, err := strconv.Atoi(strings.Replace(text, " ", "", -1))

	if err != nil {
		fmt.Println("Terjadi Kesalahan ", err.Error())
		hndl.FormInput()
		return false
	}

	if menu == 1 {
		err = hndl.SetUp()
		if err == nil {
			hndl.FormInput()
		} else {
			fmt.Println(err.Error())
		}
	} else {
		hndl.FormInput()
	}

	return true
}

func (hndl *handler) SetUp() error {
	contentCoin, errSaldo := ioutil.ReadFile("../files/import_coin.txt")
	contentItem, errItem := ioutil.ReadFile("../files/import_item.txt")

	if errSaldo != nil || errItem != nil {
		fmt.Println(errSaldo.Error())
		fmt.Println(errItem.Error())
		return errors.New("Failed get file ")
	}

	ctx := context.Background()

	_, errSaldo = hndl.saldo.Upsert(ctx, string(contentCoin))
	_, errItem = hndl.item.Upsert(ctx, string(contentItem))

	if errSaldo != nil || errItem != nil {
		fmt.Println(errSaldo.Error())
		fmt.Println(errItem.Error())
		return errors.New("Failed import data ")
	}

	return nil
}

func (hndl *handler) FormInput() bool {
	ctx := context.Background()

	fmt.Println("Coin : ")

	dataSaldo, _ := hndl.saldo.Find(ctx)
	for i, entry := range dataSaldo {
		fmt.Println(i+1, ". coin: ", entry.Value, " count: ", entry.Count)
	}

	fmt.Println("Barang Yang Dijual :")

	data, _ := hndl.item.Find(ctx)

	for i, entry := range data {
		fmt.Println(i+1, ". ", entry.Name, " harga: ", entry.Price, " Tersedia: ", entry.Stock)
	}

	fmt.Print("Input Nomor Menu : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	menu, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Terjadi Kesalahan ", err.Error())
		hndl.FormInput()
		return false
	}

	hndl.Transaction(data[menu-1])
	return true
}

func (hndl *handler) Transaction(item model.Items) bool {
	ctx := context.Background()

	fmt.Print("Input Coin : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	coin, err := strconv.Atoi(text)

	if err != nil || int64(coin) < item.Price {
		fmt.Println("Terjadi Kesalahan atau coin anda kurang")
		hndl.Transaction(item)
		return false
	}

	// check saldo
	restCoint, rest, err := hndl.saldo.CheckCoin(ctx, coin, item.Price)
	if err != nil {
		fmt.Println(err.Error())
		hndl.FormInput()
		return false
	}

	// update item
	item.Stock = item.Stock - 1
	hndl.item.Update(ctx, item)

	fmt.Println("Transaksi Berhasil : ", item.Name)
	fmt.Println("Dengan Kembalian :")
	for _, cointEntry := range restCoint {

		// update coin
		dataGetSaldo, err := hndl.saldo.Get(ctx, cointEntry.Value)

		if err == nil {
			dataGetSaldo.Count = dataGetSaldo.Count - cointEntry.Count
			hndl.saldo.Update(ctx, dataGetSaldo)
		}

		fmt.Println(cointEntry.Value, " X ", cointEntry.Count, " = ", cointEntry.Value*cointEntry.Count)
	}

	fmt.Println("TOTAL KEMBALIAN : ", rest)

	hndl.FormInput()

	return true
}
