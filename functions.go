package UntisAPI

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

func (u *User) GetTeachers() (map[int]Teacher, error) {
	response, err := u.request("getTeachers", nil)
	if err != nil {
		return nil, err
	}

	teachers := map[int]Teacher{}
	for _, data := range response.Result.([]interface{}) {

		var teacher Teacher
		err := mapstructure.Decode(data, &teacher)
		if err != nil {
			return nil, err
		}

		teachers[teacher.Id] = teacher
	}

	return teachers, nil
}

type Student struct {
	Id       int
	Key      string
	Name     string
	ForeName string
	LongName string
	Gender   string
}

func (u *User) GetStudents() (map[int]Student, error) {
	response, err := u.request("getStudents", nil)
	if err != nil {
		return nil, err
	}

	students := map[int]Student{}
	for _, data := range response.Result.([]interface{}) {

		var student Student
		err := mapstructure.Decode(data, &student)
		if err != nil {
			return nil, err
		}

		students[student.Id] = student
	}

	return students, nil
}

type Class struct {
	Id       int
	Name     string
	LongName string
	Active   bool
	Teacher1 int
	Teacher2 int
}

func (u *User) GetClasses() (map[int]Class, error) {
	response, err := u.request("getKlassen", nil)
	if err != nil {
		return nil, err
	}

	classes := map[int]Class{}
	for _, data := range response.Result.([]interface{}) {

		var class Class
		err := mapstructure.Decode(data, &class)
		if err != nil {
			return nil, err
		}

		classes[class.Id] = class
	}

	return classes, nil
}

type Subject struct {
	Id            int
	Name          string
	LongName      string
	AlternateName string
	Active        bool
}

func (u *User) GetSubjectes() (map[int]Subject, error) {
	response, err := u.request("getSubjects", nil)
	if err != nil {
		return nil, err
	}

	subjectes := map[int]Subject{}
	for _, data := range response.Result.([]interface{}) {

		var subject Subject
		err := mapstructure.Decode(data, &subject)
		if err != nil {
			return nil, err
		}

		subjectes[subject.Id] = subject
	}

	return subjectes, nil
}

type Room struct {
	Id       int
	Name     string
	LongName string
	Building string
	Active   bool
}

func (u *User) GetRooms() (map[int]Room, error) {
	response, err := u.request("getRooms", nil)
	if err != nil {
		return nil, err
	}

	rooms := map[int]Room{}
	for _, data := range response.Result.([]interface{}) {

		var room Room
		err := mapstructure.Decode(data, &room)
		if err != nil {
			return nil, err
		}

		rooms[room.Id] = room
	}

	return rooms, nil
}

type Schoolyear struct {
	Name      string
	StartDate int
	EndDate   int
}

func (u *User) GetCurrentSchoolyear() (Schoolyear, error) {
	response, err := u.request("getCurrentSchoolyear", nil)
	if err != nil {
		return Schoolyear{}, err
	}

	var year Schoolyear
	err = mapstructure.Decode(response.Result, &year)
	if err != nil {
		return Schoolyear{}, err
	}

	return year, nil
}
func (u *User) GetSchoolyears() ([]Schoolyear, error) {
	response, err := u.request("getSchoolyears", nil)
	if err != nil {
		return nil, err
	}

	var years []Schoolyear
	for _, data := range response.Result.([]interface{}) {

		var year Schoolyear
		err := mapstructure.Decode(data, &year)
		if err != nil {
			return nil, err
		}

		years = append(years, year)
	}

	return years, err
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

func (u *User) GetTimeTable(id int, idtype int, startDate int, endDate int) (map[int]Period, error) {
	param := map[string]interface{}{
		"id":        id,
		"type":      idtype,
		"startDate": startDate,
		"endDate":   endDate,
	}
	response, err := u.request("getTimetable", param)
	if err != nil {
		return nil, err
	}

	periods := map[int]Period{}
	for _, data := range response.Result.([]interface{}) {

		var period Period
		err := mapstructure.Decode(data, &period)
		if err != nil {
			return nil, err
		}

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
	return periods, nil
}

func (u *User) GetPersonId(firstname string, lastname string, isTeacher bool) (int, error) {
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
	response, err := u.request("getPersonId", param)
	return int(response.Result.(float64)), err
}
