package Untis

import (
	"github.com/mitchellh/mapstructure"
)

type Teacher struct {
	Id       int
	Name     string
	ForeName string
	LongName string
	Title    string
	Active   bool
	//Dids []interface{}
}

func (u *User) GetTeachers() map[int]Teacher {
	response := u.request("getTeachers", nil)

	teachers := map[int]Teacher{}
	for _, data := range response.Result.([]interface{}) {

		var teacher Teacher
		checkError(mapstructure.Decode(data, &teacher))
		teachers[teacher.Id] = teacher
	}

	return teachers
}

type Student struct {
	Id       int
	Key      string
	Name     string
	ForeName string
	LongName string
	Gender   string
}

func (u *User) GetStudents() map[int]Student {
	response := u.request("getStudents", nil)

	students := map[int]Student{}
	for _, data := range response.Result.([]interface{}) {

		var student Student
		checkError(mapstructure.Decode(data, &student))
		students[student.Id] = student
	}

	return students
}

type Class struct {
	Id       int
	Name     string
	LongName string
	Active   bool
	Teacher1 int
	Teacher2 int
}

func (u *User) GetClasses() map[int]Class {
	response := u.request("getKlassen", nil)

	classes := map[int]Class{}
	for _, data := range response.Result.([]interface{}) {

		var class Class
		checkError(mapstructure.Decode(data, &class))
		classes[class.Id] = class
	}

	return classes
}

type Subject struct {
	Id            int
	Name          string
	LongName      string
	AlternateName string
	Active        bool
}

func (u *User) GetSubjectes() map[int]Subject {
	response := u.request("getSubjects", nil)

	subjectes := map[int]Subject{}
	for _, data := range response.Result.([]interface{}) {

		var subject Subject
		checkError(mapstructure.Decode(data, &subject))
		subjectes[subject.Id] = subject
	}

	return subjectes
}

type Room struct {
	Id       int
	Name     string
	LongName string
	Building string
	Active   bool
}

func (u *User) GetRooms() map[int]Room {
	response := u.request("getRooms", nil)

	rooms := map[int]Room{}
	for _, data := range response.Result.([]interface{}) {

		var room Room
		checkError(mapstructure.Decode(data, &room))
		rooms[room.Id] = room
	}

	return rooms
}

type Schoolyear struct {
	Name      string
	StartDate int
	EndDate   int
}

func (u *User) GetCurrentSchoolyear() Schoolyear {
	response := u.request("getCurrentSchoolyear", nil)

	var year Schoolyear
	checkError(mapstructure.Decode(response.Result, &year))

	return year
}
func (u *User) GetSchoolyears() []Schoolyear {
	response := u.request("getSchoolyears", nil)

	var years []Schoolyear
	for _, data := range response.Result.([]interface{}) {

		var year Schoolyear
		checkError(mapstructure.Decode(data, &year))
		years = append(years, year)
	}

	return years
}

type Period struct {
	StartTime    int
	ActivityType string
	Id           int
	Date         int
	EndTime      int
	Classes      []int
	Subject      []int
	Teacher      []int
	Rooms        []int
}

func (u *User) GetTimeTable(id int, idtype int, startDate int, endDate int) map[int]Period {
	param := map[string]interface{}{
		"id":        id,
		"type":      idtype,
		"startDate": startDate,
		"endDate":   endDate,
	}
	response := u.request("getTimetable", param)

	periods := map[int]Period{}
	for _, data := range response.Result.([]interface{}) {

		var period Period
		checkError(mapstructure.Decode(data, &period))

		dataMap := data.(map[string]interface{})
		for _, klasse := range dataMap["kl"].([]interface{}) {
			period.Classes = append(period.Classes, int(klasse.(map[string]interface{})["id"].(float64)))
		}
		for _, student := range dataMap["su"].([]interface{}) {
			period.Subject = append(period.Subject, int(student.(map[string]interface{})["id"].(float64)))
		}
		for _, room := range dataMap["ro"].([]interface{}) {
			period.Rooms = append(period.Rooms, int(room.(map[string]interface{})["id"].(float64)))
		}
		for _, teacher := range dataMap["te"].([]interface{}) {
			period.Teacher = append(period.Teacher, int(teacher.(map[string]interface{})["id"].(float64)))
		}

		periods[period.Id] = period

	}
	return periods
}

func (u *User) GetPersonId(firstname string, lastname string, isTeacher bool) int {
	param := map[string]interface{}{
		"fn":  firstname,
		"sn":  lastname,
		"dob": "0",
	}
	if isTeacher {
		param["type"] = "2"
	} else {
		param["type"] = "5"
	}
	response := u.request("getPersonId", param)
	return int(response.Result.(float64))
}
