# SOLID
#### SOLID principles are object-oriented design concepts relevant to software development. SOLID is an acronym for five other class-design principles: **Single Responsibility Principle**, **Open-Closed Principle**, **Liskov Substitution Principle**, **Interface Segregation Principle**, and **Dependency Inversion Principle**.
***
SOLID tamoyillari dasturiy ta'minotni ishlab chiqishga tegishli ob'ektga yo'naltirilgan dizayn tushunchalari. SOLID - bu sinfning boshqa beshta dizayn tamoyillarining qisqartmasi: **Yagona javobgarlik printsipi**, **ochiq-yopiq printsip**, **Liskov almashtirish printsipi**, **interfeyslarni ajratish printsipi** va **qaramlikni inversiya qilish printsipi**.

## **S - Single Responsibility Principle** - Yagona javobgarlik printsipi
Each class should be responsible for a single part or functionality of the system. </br>
Har bir sinf tizimning bir qismi yoki funksionalligi uchun javobgar bo'lishi kerak.


## **O - Open-Closed Principle** - ochiq-yopiq printsipi
Software components should be open for extension, but not for modification. </br>
Dasturiy ta'minot komponentlari kengaytirish uchun ochiq bo'lishi kerak, lekin o'zgartirish uchun emas.

## **L - Liskov Substitution Principle** - Liskov almashtirish printsipi
Objects of a superclass should be replaceable with objects of its subclasses without breaking the system. </br>
Supersinf ob'ektlari tizimni buzmasdan uning kichik sinflari ob'ektlari bilan almashtirilishi kerak.

## **I - Interface Segregation Principle** - interfeyslarni ajratish printsipi
No client should be forced to depend on methods that it does not use. </br>
Hech bir mijoz o'zi ishlatmaydigan usullarga qaram bo'lishga majburlanmasligi kerak.

## **D - Dependency Inversion Principle** - qaramlikni(bog'liqlik) inversiya qilish printsipi
High-level modules should not depend on low-level modules, both should depend on abstractions. </br>
Yuqori darajadagi modullar past darajadagi modullarga bog'liq bo'lmasligi kerak, ikkalasi ham abstraktsiyalarga bog'liq bo'lishi kerak.

***************************************************************************************************
To'g'ri loyhalashtirilgan dasturlarda uning tizimi qanchalik katta va murakkab bo'lmasin, Dasturga o'zgartirish kiritish ancha oson kechadi! Sababi dasturchi bu dasturga o'zgartish uchun Tizimining barcha modullari haqida to'liq ma'lumotga ega bo'lishi talab etilmaydi! 
Bundan tashqai kiritilgan o'zgarish Dasturing boshqa modullariga ta'sir etmaydi. Sababi modullar orasidagi bog'liqlik (qaramlilik kam bo'ladi) va har bir klass alohida vazifani bajarish uchun mas'ul bo'lganligi sababli qaytariluvhi kodningham kiritilishi talab qilinmaydi.

