package main

import (
   "fmt"
)

type Patient struct {
	first string
	last string
}

type Doctor struct {
	first string
	last string
	schedule []Patient 	//size 16, presumably for 8 hr workday with half-hour chunks
}

var empty Patient

func main() {
	p1 := Patient{"Johniffer", "Doe"}
	p2 := Patient{"Janeld", "Doe"}
	doc := Doctor{"No", "Body", make([]Patient, 16)}

	fmt.Println("A doctor's default schedule.\n")
	doc.PrintSchedule()
	doc.PrintAvailableSchedule()

	fmt.Printf("Scheduling %q\n", p1)
	doc.AddPatient(p1, 15)
	doc.PrintSchedule()

	fmt.Printf("Scheduling %q\n", p2)
	doc.AddPatient(p2, 2)
	doc.PrintSchedule()

	fmt.Printf("%q is a bit of a hypochondriac.\n", p1)
	for i := 0; i < 16; i++ {
		if i%2 != 0 { 
			doc.AddPatient(p1, i)
		}
	}

	doc.PrintSchedule()
	doc.PrintAvailableSchedule()
}

//while the paramter is called "time", it should actually be the position of the array into which the Patient object will be placed.
//the correct array position should be easily determined from the output of the `PrintAvailableSchedule` method
func (d *Doctor) AddPatient(p Patient, time int) {
	d.schedule[time] = p
}

func (d *Doctor) RemovePatient(time int) {
	d.schedule[time] = empty
}

func (p Patient) String() string {
	return p.first + " " + p.last
}

func (d Doctor) String() string {
	return fmt.Sprintf("Dr. %s %s", d.first, d.last)
}

func (d Doctor) PrintSchedule() {
	fmt.Printf("%s's Appointments :\n", d)
	for i,v := range d.schedule {
		switch i {
		case 0:
			fmt.Print("\t08:00 || ")
		case 1:
			fmt.Print("\t08:30 || ")
		case 2:
			fmt.Print("\t09:00 || ")
		case 3:
			fmt.Print("\t09:30 || ")
		case 4:
			fmt.Print("\t10:00 || ")
		case 5:
			fmt.Print("\t10:30 || ")
		case 6:
			fmt.Print("\t11:00 || ")
		case 7:
			fmt.Print("\t11:30 || ")
		case 8:
			fmt.Print("\t12:00 || ")
		case 9:
			fmt.Print("\t12:30 || ")
		case 10:
			fmt.Print("\t13:00 || ")
		case 11:
			fmt.Print("\t13:30 || ")
		case 12:
			fmt.Print("\t14:00 || ")
		case 13:
			fmt.Print("\t14:30 || ")
		case 14:
			fmt.Print("\t15:00 || ")
		case 15:
			fmt.Print("\t15:30 || ")
		}
		fmt.Println(v)
	}
	fmt.Println()
}

func (d Doctor) PrintAvailableSchedule() {
	fmt.Printf("%s's Available Schedule :\n", d)
	for i,v := range d.schedule {
		if (v.first == "" && v.last == "") {
			switch i {
			case 0:
				fmt.Printf("\t%d || 08:00\n", i)
			case 1:
				fmt.Printf("\t%d || 08:30\n", i)
			case 2:
				fmt.Printf("\t%d || 09:00\n", i)
			case 3:
				fmt.Printf("\t%d || 09:30\n", i)
			case 4:
				fmt.Printf("\t%d || 10:00\n", i)
			case 5:
				fmt.Printf("\t%d || 10:30\n", i)
			case 6:
				fmt.Printf("\t%d || 11:00\n", i)
			case 7:
				fmt.Printf("\t%d || 11:30\n", i)
			case 8:
				fmt.Printf("\t%d || 12:00\n", i)
			case 9:
				fmt.Printf("\t%d || 12:30\n", i)
			case 10:
				fmt.Printf("\t%d || 13:00\n", i)
			case 11:
				fmt.Printf("\t%d || 13:30\n", i)
			case 12:
				fmt.Printf("\t%d || 14:00\n", i)
			case 13:
				fmt.Printf("\t%d || 14:30\n", i)
			case 14:
				fmt.Printf("\t%d || 15:00\n", i)
			case 15:
				fmt.Printf("\t%d || 15:30\n", i)
			}
		}
	}
	fmt.Println()		//just for a bit of clearance
}