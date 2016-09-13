
### chhs

Take the downloaded chhs json file and convert it into something that works for elastic.

Meaning we need to grab the column names of the original csv format which was converted into json and map the column names onto the json file.  Originally each row is just a line with the data in it, and we need to map the column names onto the rows.

 * json01.go reduces the big file down to the little file so it is manageable.
