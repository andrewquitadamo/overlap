package main
import "os"
import "fmt"
import "bufio"
import "strings"

func main() {
	var id_index  map[string][]int
	id_index = make(map[string][]int)
	var id_arr []string
    files := os.Args[1:]
    for index := range files {
		filename := files[index]
		f, err := os.Open(filename)
    	if err != nil {
        	fmt.Println(err)
        	return
    	}
    	defer f.Close()
    	r := bufio.NewReader(f)
    	line, err := r.ReadString('\n')
		line = strings.Trim(line,"\n") 
		larr := strings.Split(line,"\t")
		for samp := range larr {
			_, exists := id_index[larr[samp]]
			id_index[larr[samp]] = append(id_index[larr[samp]],samp) 
			if !exists {
				id_arr = append(id_arr,larr[samp])
			}
		}
	}
    for index := range files {
		var pos []int
		for k := range id_arr {
			if len(id_index[id_arr[k]])==len(files) {
				pos = append(pos,id_index[id_arr[k]][index])
			}
		}
		fmt.Println(pos)

    	filename := files[index]
    	f, err := os.Open(filename)
    	if err != nil {
        	fmt.Println(err)
        	return
    	}
    	defer f.Close()
		scanner := bufio.NewScanner(f)
		fo, err := os.Create(filename + ".out")
		w := bufio.NewWriter(fo)
		for scanner.Scan() {
			var line []string
			larr := strings.Split(scanner.Text(),"\t")
			for j := range pos {
				line = append(line,larr[j])
			}
			//fmt.Println(strings.Join(line[:],"\t"))
			w.WriteString(strings.Join(line[:],"\t") + "\n")
		}
		w.Flush()
	}
}
