# Building an HTTP based API With Go

Course material for GopherCon 2016 workshop

## Contents

* slides/ contains the presentation slides in present format.
* exercises/ contains the workshop exercises and solutions.
* workshop/ contains a complete implementation of an service written in goa for reference.

### Slides

The slides are written in [present](https://godoc.org/golang.org/x/tools/present). The directory
contains the slide sources as well as the example code and images.

### Exercises

Each exercise is contained in a directory prefix with the number of the corresponding slide deck.
The directories contain a README describing the goal of the exercise as well as a `solution/`
directory containing a possible implementation.

Some exercises build on previous ones, use the provided solution if you were not able to complete
the previous implementation.

### Workshop

The GoWorkshop service is implemented using [https://github.com/goadesign/goa](goa). It showcases a
number of advanced topics such as implementing JWT and provides a good reference for future work.
