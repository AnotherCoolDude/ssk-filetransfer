package proadclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// PostTask creates a new Task
func PostTask(task *Task) error {
	requestBody, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error marshaling task: %s\n", err)
		return err
	}
	req := makePOSTRequest("tasks", bytes.NewBuffer(requestBody))
	resp := Client.Do(req)
	if resp.StatusCode != 200 {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return fmt.Errorf("error creating task (code %d):\n%+v", resp.StatusCode, string(respBytes))
	}
	return nil
}
