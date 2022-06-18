### Description

Simple server with CRUD operations. Built using Gorilla/Mux package. 

### Challenges and lessons

During the development I have run into and learned more about the following:

- Capitalize variables in struct so that json library can access it. In Go, you need to capitalize in order to export the variable (aka make it accessible for outside functions). 
- Pointers. This is a whole new concept for me. I have researched and learned more about why do we need pointers, deferencing, as well as how the code is run in stacks and heaps. Advantages / Disadvantage from the efficiency point of view for using pointers instead of copying the variables