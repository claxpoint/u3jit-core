... ...
NUM_COL = 5
NUM_ROW = 2

.code
main PROC
   mov esi, OFFSET ary2D
   mov eax, 31h 
   call Search2DAry
; See eax for search result
   exit
main ENDP

;------------------------------------------------------------
Search2DAry PROC
; Receives: EAX, a byte value to search a 2-dimensional array
;           ESI, an address to the 2-dimensional array
; Returns: EAX, 1 if found, 0 if not found
;------------------------------------------------------------
   mov  ecx,NUM_ROW        ; outer loop count
... ...
   mov  ecx,NUM_COL        ; inner loop counter
... ...
