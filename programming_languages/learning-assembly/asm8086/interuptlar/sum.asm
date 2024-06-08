ORG 100H

    MOV AH, 9
    LEA DX, msg1
    INT 21h
    
    MOV AH, 1
    INT 21h
    MOV BH, AL

    MOV AH, 9
    LEA DX, msg2
    INT 21h 
    
    MOV AH, 1
    INT 21h
    MOV BL, AL
    
    ADD BH, BL
    SUB BH, 48
    
    MOV AH, 2
    MOV DL, BH
    INT 21h

RET 
    msg1 DB, 13, 10, "enter first number:  $"
    msg2 DB, 13, 10, "enter second number: $"

END