package event

import (
	"email-service-event-driven/service"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

type ProductCreatedEvent struct {
	NamaProduk   string  `json:"nama_produk,omitempty"`
	Harga        *int    `json:"harga,omitempty"`
	Kategori     string  `json:"kategori,omitempty"`
	Deskripsi    *string `json:"deskripsi,omitempty"`
	Stok         *int    `json:"stok,omitempty"`
	LastModified int64   `json:"last_modified,omitempty"`
	Service      service.EmailService
}

func (pc ProductCreatedEvent) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	err := json.Unmarshal(m.Body, &pc)
	if err != nil {
		return err //500 Intenral Server Error
	}
	log.Printf("received data: %#v\n", pc)
	message := `Successfully Created Product<br>
	<ul>
		<li>Product Name:<b>` + pc.NamaProduk + `</b></li>
		<li>Harga:<b>` + strconv.Itoa(*pc.Harga) + `</b></li>
		<li>Kategori:<b>` + pc.Kategori + `</b></li>
		<li>Deskripsi:<b>` + *pc.Deskripsi + `</b></li>
		<li>Stok:<b>` + strconv.Itoa(*pc.Stok) + `</b></li>
		<li>Terakhir diubah:<b>` + time.UnixMilli(pc.LastModified).String() + `</b></li>
	</ul>
	`
	err = pc.Service.SendEmail(message)
	if err != nil {
		panic(err) //500 internal, tidak bisa kirim email
	}
	return nil
}
