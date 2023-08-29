package usecase

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/test-jagaad/internal/domain"
	"github.com/test-jagaad/internal/entity"
)

type (
	UserUc interface {
		Fetch() (resp entity.FetchUserResp, err error)
		Search(tags []string) (resp []string, err error)
	}

	userUc struct {
		mockyDom domain.MockyDom
	}
)

func NewUserUc(mockyDom domain.MockyDom) UserUc {
	return &userUc{mockyDom}
}

func (u *userUc) Fetch() (resp entity.FetchUserResp, err error) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	chanResults := make(chan entity.Wrapper)
	go func() {
		defer wg.Done()
		users1, errDom := u.mockyDom.GetMocky1()
		chanResults <- entity.Wrapper{Name: "users1", Result: users1, Err: errDom}
	}()
	go func() {
		defer wg.Done()
		users2, errDom := u.mockyDom.GetMocky2()
		chanResults <- entity.Wrapper{Name: "users2", Result: users2, Err: errDom}
	}()
	go func() {
		wg.Wait()
		close(chanResults)
	}()

	csvRecords := []entity.CsvRecord{}
	for res := range chanResults {
		if res.Err != nil {
			resp.Details = append(resp.Details, entity.FetchUserRespDetail{
				Name: res.Name,
				Err:  res.Err,
			})
			continue
		}

		if res.Result != nil {
			for _, user := range res.Result {
				friendNames := make([]string, len(user.Friends))
				for k, friend := range user.Friends {
					friendNames[k] = friend.Name
				}
				csvRecords = append(csvRecords, entity.CsvRecord{
					Data: []string{user.Balance, strings.Join(user.Tags, ","), strings.Join(friendNames, ",")},
				})
			}
			resp.Details = append(resp.Details, entity.FetchUserRespDetail{
				Name: res.Name,
			})
		}
	}

	resp.Filename = entity.UserFilename
	err = u.saveToCsv(entity.UserFilename, csvRecords)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (u *userUc) Search(tags []string) (resp []string, err error) {
	file, err := os.Open(fmt.Sprintf("%stmp/%s.csv", os.Getenv("TMP_DIR"), entity.UserFilename))
	if err != nil {
		return resp, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}

		mapTagFreq := make(map[string]int)
		for _, tag := range tags {
			mapTagFreq[tag]++
		}

		userTags := strings.Split(record[1], ",")
		mapUserTagFreq := make(map[string]int)
		for _, tag := range userTags {
			mapUserTagFreq[tag]++
		}

		found := true
		for tag := range mapUserTagFreq {
			if val, ok := mapTagFreq[tag]; !ok || mapTagFreq[tag] != val {
				found = false
				break
			}
		}

		if found {
			resp = append(resp, fmt.Sprintf("Balance: %s, Friends: %s", record[0], record[2]))
		}
	}

	return resp, nil
}

func (u *userUc) saveToCsv(filename string, csvRecords []entity.CsvRecord) error {
	file, err := os.Create(fmt.Sprintf("%stmp/%s.csv", os.Getenv("TMP_DIR"), entity.UserFilename))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := range csvRecords {
		writer.Write(csvRecords[i].Data)
	}

	return nil
}
