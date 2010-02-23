//	TODO:	check both synchronous and asynchronous channels are supported
//	TODO:	Read() and Write() interface support
//	TODO:	write tests

package vm

type IOController []chan *Stream

func (ioc *IOController) realloc(length int) (c []chan *Stream) {
	c = make([]chan *Stream, length)
	copy(c, *ioc)
	*ioc = c
	return
}

func (ioc *IOController) Open(c chan *Stream) {
	starting_length := ioc.Len()
	ioc.realloc(starting_length + 1)
	ioc.Set(starting_length, c)
}

func (ioc *IOController) Clone() *IOController {
	c := new(IOController)
	c.realloc(ioc.Len())
	copy(*c, *ioc)
	return c
}

func (ioc *IOController) Init()									{ ioc.realloc(0) }
func (ioc *IOController) Len() int								{ return len(*ioc) }
func (ioc *IOController) Cap() int								{ return cap(*ioc) }
func (ioc *IOController) At(i int) chan *Stream					{ return (*ioc)[i] }
func (ioc *IOController) Set(i int, c chan *Stream)				{ (*ioc)[i] = c }
func (ioc *IOController) Close(i int)							{ close(ioc.At(i)) }
func (ioc *IOController) CloseAll()								{ for i, _ := range *ioc { ioc.Close(i) } }
func (ioc *IOController) Send(i int, x *Stream)					{ go func() { ioc.At(i) <- x.Clone() }() }
func (ioc *IOController) Receive(i int) *Stream					{ return <-ioc.At(i) }
func (ioc *IOController) Copy(i, j int)							{ ioc.At(i)<- <-ioc.At(j) }

//func (ioc *IOController) Do(f func(x *Stream))				{ for _, e := range *ioc { f(<-e) } }