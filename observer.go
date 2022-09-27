package main

import (
	"fmt"
)

func main() {
	firstJob := NewJob("Java Developer")

	firstPerson := &Person{pName: "Jones"}
	secondPerson := &Person{pName: "Maria"}

	firstJob.subscribe(firstPerson)
	firstJob.subscribe(secondPerson)

	firstJob.updateAvailability()

	firstJob.unsubscribe(firstPerson)

	firstJob.updateAvailability()
}

type Observer interface {
	update(string)
	getID() string
}

type Observable interface {
	subscribe()
	unsubscribe()
	notify()
}

type Person struct {
	pName string
}

type Job struct {
	observerList []Observer
	name         string
	availability bool
}

func NewJob(name string) *Job {
	return &Job{
		name: name,
	}
}

func (j *Job) updateAvailability() {
	fmt.Printf("Job %s is available now\n", j.name)
	j.availability = true
	j.notify()
}

func (j *Job) subscribe(o Observer) {
	j.observerList = append(j.observerList, o)
}

func (j *Job) unsubscribe(o Observer) {
	j.observerList = removeList(j.observerList, o)
}

func (j *Job) notify() {
	for _, observer := range j.observerList {
		observer.update(j.name)
	}
}

func removeList(observerList []Observer, ObserverToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if ObserverToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (p *Person) update(jobName string) {
	fmt.Printf("Sending email to %s for job %s\n", p.pName, jobName)
}

func (p *Person) getID() string {
	return p.pName
}
