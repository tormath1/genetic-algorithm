## Simple genetic algorithm in Go

To run this algo, you'll need to have Go installed on your machine:

```shell
$ go version
go version go1.10.3 linux/amd64
```

You can now change the value for `target` and run the script (`$ go run main.go`). You can also tune the parameters (mutation prob, etc):

```shell
$ ## the value to converge on is `easy`
$ go run main.go 
2018/06/11 12:04:50 [[d t w e] [h g c i] [z r b h] [v b g g] [r k h o] [c i o k] [e y p r] [z u c d] [y k p v] [e g t f] [u f v d] [j v a c] [z g l j] [k s i f] [t g y a] [n i p n] [t i p a] [p f r i] [i s p w] [d x y j] [p b i c] [k m u y] [b k m l] [y j j n] [l k x r] [o u i j] [x c b x] [y o x u] [g g v w] [y v j w] [x o c u] [k y f i] [h y n t] [t j c h] [g t n n] [d m f i] [j k u c] [w l a r] [i c r f] [o a i d] [c d c a] [d f x v] [j b c d] [z y s q] [e q n y] [u z a l] [q o h p] [h m a t] [z w m g] [l r l i] [c b x s] [p v r y] [a q i h] [v w x i] [u b k g] [f w o c] [a w b h] [m i j j] [g e g k] [v p c i] [x i a q] [w y i h] [o q l t] [t f o n] [o x m z] [o i k y] [q l c w] [g e a y] [o i d a] [k g o k] [k i j x] [f f g o] [i n j g] [g p v w] [f h x o] [n y i i] [m c g h] [l v a r] [a o x e] [o p p t] [a w k z] [r g e g] [b c y w] [c g v k] [p k e g] [i x l h] [r k i v] [a u b a] [n e j a] [v h n e] [h w u i] [u p w s] [x h t k] [a i g w] [q u u m] [c p g x] [a t c k] [v r m m] [l l w f]]
2018/06/11 12:04:50 [[e j a y] [e a a y] [e a a y] [e a a y] [b a a y] [e a a y] [e a n y] [e a a y] [e a n y] [e a a y] [e a a s] [e a a q] [e a a b] [e a a y] [e a a y] [e g s y] [e a a y] [e a a y] [e a a y] [e g a y] [e a a y] [e a a y] [e a s y] [e a a y] [b a a y] [e a n y] [e x a y] [e g s y] [e a a y] [e a a y] [e a a y] [e a n y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e x a y] [e a a y] [e a a y] [e a a y] [e a a y] [e j a y] [e a a y] [e a a y] [e a a q] [e a u y] [e a n y] [e a a y] [e a a y] [e a a y] [e a a y] [e g s y] [e a a y] [e a m y] [e a a y] [e a a y] [e a a y] [e a n y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a u y] [e a a y] [e a n y] [e a a y] [e a n y] [e a a y] [e a a y] [e k a y] [e a n y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y] [e a a y]] 18
```
