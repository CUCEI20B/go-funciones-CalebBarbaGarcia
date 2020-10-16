package main

func Burbuja(s []int64)  {
	var i int
	var j int
	i = 0
	
	for (i < len(s)){
		j = i
		for (j < len(s)){
			if (s[j] < s[i]){
				aux := s[j]
				s[j] = s[i]
				s[i] = aux
			}

			j = j + 1
		}
		i = i +1
	}
}

func main()  {
	
}