package main
import "fmt"
const MAX_SIZE = 5
type Queue struct{
	 a[MAX_SIZE] int
	l int 
	r int 
	sz int 
}
func NewQueue() *Queue{
	q:= &Queue{
		l:0,
		r:-1,
		sz:0,
	}
	return q
}
func (q *Queue) Enqueue(x int ){  //x=1, x=6 , x=7 , //x=9 , x=16
	if q.sz==MAX_SIZE {    //sz=0, sz=1 ,sz =2, //sz=3 , sz=4 , sz=5 
		fmt.Println("Queue is full")
		return
	}
	q.r++ //r=-1+1=0 , r=0+1=1 , r=1+1=2 //r=2+1=3 , r=3+1=4 , r=4+1=5
	if q.r==MAX_SIZE{
		q.r=0

	}
	q.a[q.r]=x //a[0]=1 ,a[1]=6 ,a[2]=7, //a[3]=9 , a[4]=16 , a[0]=16
q.sz++

}
func (q *Queue) Dequeue(){
	if q.sz==0{
		fmt.Println("Queue is empty")
		return
	}
	q.l++
	if q.l==MAX_SIZE {
	q.l=0
	}
	q.sz--
}
func (q *Queue) Front() int {
	if q.sz==0 {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.a[q.l]
}
func (q *Queue) Size() int {
	return q.sz
}

func main (){
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(9)
	q.Enqueue(16)
	q.Enqueue(16) // "Queue is full"
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	fmt.Print(q.Front(), " ")
	q.Dequeue()	
	q.Dequeue()
	q.Dequeue()

	fmt.Print(q.Front(), " ")
	fmt.Println("\nSize of queue:", q.Size())
	q.Enqueue(20)
	fmt.Print(q.Front(), " ")
	fmt.Println("\nSize of queue after enqueue:", q.Size())


}