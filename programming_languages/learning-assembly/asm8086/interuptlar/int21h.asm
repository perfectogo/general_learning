;input
ORG 100H
    ; if ah = 1
    MOV AH, 1
    ; write one value to AL
    INT 21h
RET

;output 
ORG 100H
    ; if ah = 2
    MOV AH, 2
    MOV DL, 'A'
    ; read DL's value 
    INT 21h
RET

ORG 100H

    MOV AH, 9
    MOV BX, OFFSET word
    INT 21h
RET 

    word db "Hello, World $"
END

ORG 100H
    MOV AH, 0Ah
    MOV DX, OFFSET word
    INT 21h

RET
    word DB 21, ?, 21
END

