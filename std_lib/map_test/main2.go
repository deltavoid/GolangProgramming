package main

import "fmt"




func new_map1() map[string]int {

    m := make(map[string]int)

    m["k1"] = 1

    fmt.Printf("new_map1, m: %p\n", &m)
    return m
}


func new_map2() map[string]int {

    m := map[string]int{"foo": 1, "bar": 2}

    fmt.Printf("new_map2: m: %p\n", &m)
    return m
}

func new_map3() *map[string]int {

    m := map[string]int{"foo": 1, "bar": 2}

    fmt.Printf("new_map3: m: %p\n", &m)
    return &m
}

func display_map1(m map[string]int) {

    fmt.Printf("display_map1, m: %p\n", &m)

    for key, val := range m {

        fmt.Printf("%s: %d\n", key, val)

    }

}


func display_map2(m *map[string]int) {

    fmt.Printf("display_map2, m: %p\n", m)

    for key, val := range *m {

        fmt.Printf("%s: %d\n", key, val)

    }

}


func update_map(m map[string]int) {

    m["hello"] = 1

}


func main() {


    m1 := new_map1()
    fmt.Printf("m1: %p\n", &m1)
    display_map1(m1)
    display_map2(&m1)
    fmt.Println()


    m2 := new_map2()
    fmt.Printf("m2: %p\n", &m2)
    display_map1(m2)
    display_map2(&m2)
    fmt.Println()
    

    m3_p := new_map3()
    fmt.Printf("m3: %p\n", m3_p)
    display_map1(*m3_p)
    display_map2(m3_p)
    fmt.Println()

    
    m4 := new_map1()
    fmt.Printf("m4: %p\n", m4)
    update_map(m4)
    display_map1(m4)


}
