
### chhs

Take the downloaded chhs json file and convert it into something that works for elastic.

Meaning we need to grab the column names of the original csv format which was converted into json and map the column names onto the json file.  Originally each row is just a line with the data in it, and we need to map the column names onto the rows.

With further research the easiest way to do this is to grab the column names from the csv file and grab the json data from the json file and then merge the two to create a json document that is ready to go into elastic, bleve, noms, etc...

Trying to read the column data out of the json file is not a good idea because the column names that come out of the json file which is stored in a map is not guaranteed every time to line up with the actual json column data.

 * json01.go reduces the big json file down to the little file so it is manageable.
 * csv01.go reduces the big csv down to the little file so it is manageable.
 * csv02.go contains a method to grab the column names.
