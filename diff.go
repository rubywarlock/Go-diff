package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

type RSpecList struct {
    // Name string
    Data []string
}

func main() {
    patches := [2]string{"./list_two.json", "./list_ones.json"}
    var lists [len(patches)]RSpecList
    var ptr_lists []*RSpecList

    for i, patch := range patches {
        data, err := ioutil.ReadFile(patch)
        var list RSpecList

        if err != nil {
            log.Fatal("Error file read: ", err)
        }

        err = json.Unmarshal(data, &list)

        if err != nil {
            log.Fatal("Error during Unmarshal(): ", err)
        } else {
            lists[i] = list
            ptr := &lists[i]
            ptr_lists = append(ptr_lists, &(*ptr))
        }
    }

    var missings []string

    for i := 0; i < len(lists); i++ {
        observable_ptr := ptr_lists[0]
        observable := *observable_ptr
        ptr_lists = ptr_lists[1:]

        for _, source_ptr := range ptr_lists {
            source := *source_ptr

            for observable_i, observable_value := range observable.Data {
                found := false

                for _, source_value := range source.Data {
                    if eq := (observable_value == source_value); eq == true {
                        found = true
                    }
                }

                if !found {
                    value := fmt.Sprintf("%d; %s", observable_i, observable_value)
                    missings = append(missings, value)
                }
            }

            missings = append(missings, "-------------------------------------------------------------------------------------------------")
        }

        ptr_lists = append(ptr_lists, observable_ptr)
    }

    if len(missings) > 0 {
        for _, value := range missings {
            fmt.Println(value)
        }
    } else {
        fmt.Println("not found missings")
    }
}
