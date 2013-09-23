package goption

import (
    "testing"
)

func sampleOption() (c *Option) {
    c = NewOption(map[string]interface{} {
        "name": "Tom",
        "color": "Brown",
        "height": 0.25,
        "age": 3,
        "male": true,
    })

    return
}

func TestMain(t *testing.T) {
    c := sampleOption()

    if name := c.MustGetString("name"); name != "Tom" {
        t.Errorf("MustGetString error for key: %s", "name")
    }

    if weight, ok := c.GetInt("weight"); weight != 0 || ok != false {
        t.Errorf("Key should not exist: %s", "weight")
    }

    if height, ok := c.GetFloat32("height"); height != 0.25 || ok != true {
        t.Errorf("GetFloat32 error for key: %s", "height")
    }

    if age, ok := c.GetInt("age"); age != 3 || ok != true {
        t.Errorf("GetInt error for key: %s", "age")
    }

    if male, ok := c.GetBool("male"); ok != true || male != true {
        t.Errorf("GetBool error for key: %s", "male")
    }

    if height, ok := c.GetInt("height"); ok != true || height != 0 {
        t.Errorf("GetInt error for key: %s", "height")
    }
}

func TestMerge(t *testing.T) {
    c := sampleOption()

    c.MergeMap(map[string]interface{} {
        "name": "Jack",
        "weight": 5.5,
        "male": false,
    })

    if name := c.MustGetString("name"); name != "Jack" {
        t.Errorf("MustGetString error for key: %s", "name")
    }

    if weight, ok := c.GetFloat32("weight"); weight != 5.5 || ok != true {
        t.Errorf("GetFloat32 error for key: %s", "weight")
    }

    if male, ok := c.GetBool("male"); male != false || ok != true {
        t.Errorf("GetBool error for key: %s", "male")
    }

    if height, ok := c.GetFloat32("height"); height != 0.25 || ok != true {
        t.Errorf("GetFloat32 error for key: %s", "height")
    }
}