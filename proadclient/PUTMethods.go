package proadclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// PutTask creates a new Task
func PutTask(task *Task) error {
	if task.Urno == 0 {
		return errors.New("valid task urno required")
	}
	requestBody, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error marshaling task: %s\n", err)
		return err
	}
	req := makePUTRequest("tasks", task.Urno, bytes.NewBuffer(requestBody))
	resp := Client.Do(req)
	if resp.StatusCode != 200 {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return fmt.Errorf("error updating task (code %d):\n%+v", resp.StatusCode, string(respBytes))
	}
	return nil
}
