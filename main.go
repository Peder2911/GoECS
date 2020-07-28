
package main

// import "fmt"
import (
   "image"
   "image/png"
   "image/color"
   "os"
   "fmt"
)

type Point struct {
   x float32
   y float32
}

func (a Point) add(b Point) Point {
   return Point{
      x: a.x+b.x,
      y: a.y+b.y,
   }
}

func (p Point) draw(img *image.RGBA, size int, color color.Color) {
   for xi := int(p.x)-size ; xi < int(p.x)+size ; xi++ {
      for yi := int(p.y)-size ; yi< int(p.y)+size ; yi++ {
         img.Set(xi,yi,color)
      }
   }
}

type State struct{
   Position map[int]Point
   Systems []func()
}

func (s State) iterate(){
   for _,fn := range s.Systems {
      fn()
   }
}

func moverFactory(s State) func() {
   return func(){
      s.Position[0] = s.Position[0].add(Point{10,10})
   }
}

func initializeState() State {
   s := State {
      Position: make(map[int]Point),
   }
   s.Systems = append(s.Systems,moverFactory(s))
   s.Position[0] = Point{10,10}
   s.Position[1] = Point{310,310}
   return s
}

func main(){
   ul := image.Point{0,0}
   lr := image.Point{320,320}
   white := color.RGBA{255,255,255,0xff}
   gray := color.RGBA{100,100,100,0xff}
   state := initializeState()

   for i:=0; i<10; i++ {
      img := image.NewRGBA(image.Rectangle{ul,lr})
      for x:=0;x<320;x++{
         for y:=0;y<320;y++{
            img.Set(x,y,gray)
            fmt.Printf("x:%v y:%v\n",x,y)
         }
      }
      state.iterate()
      fmt.Println(state)
      for _,p := range state.Position {
         p.draw(img,10,white)
      }
      fname := fmt.Sprintf("out/img_%v.png",i)
      f, _ := os.Create(fname)
      png.Encode(f,img)
   }

}
