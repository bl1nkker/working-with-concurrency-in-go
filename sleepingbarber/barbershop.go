package sleepingbarber

import (
	"time"

	"github.com/fatih/color"
)

type Barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (b *Barbershop) AddBarber(barber string) {
	b.NumberOfBarbers++

	go func() {
		// Initially barber is awake
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)
		for {
			// If there are no clients, the barber goes to sleep
			if len(b.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap...", barber)
				isSleeping = true
			}
			// The second parameter, is primarily used to determine whether the channel is open or closed. So, when we close our barbershop
			// we will close the channel and all barbers will go home
			client, shopOpen := <-b.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s", client, barber)
					isSleeping = true
				}
				// cut hair
				b.CutHair(barber, client)
			} else {
				// shop is closed. send this barber home
				b.SendBarberHome(barber)
				//  and close this go routine
				return
			}
		}
	}()
}

func (b *Barbershop) CutHair(barber string, client string) {
	color.Magenta("%s is cutting %s's hair", barber, client)
	time.Sleep(b.HairCutDuration)
	color.Magenta("%s is finished cutting %s's hair", barber, client)
}

func (b *Barbershop) SendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	b.BarbersDoneChan <- true
}

func (b *Barbershop) CloseShopForDay() {
	color.Green("Closing shop for the day...")
	close(b.ClientsChan)
	b.Open = false

	// wait until all the barbers are done
	for a := 1; a <= b.NumberOfBarbers; a++ {
		// that will block until every barber inside NumberOfBarbers will send some data (expected to be true)
		<-b.BarbersDoneChan
	}

	close(b.BarbersDoneChan)
	color.Green("The barbershop is now closed for the day and everyone hase gone home.")
	color.Green("--------------------------")
}
