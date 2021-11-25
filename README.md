# gencsv
gensv - file generator

Csv file generator. Input file example:
```json
{
  "filename": "report.csv",
  "rows": 1000,
  "header": "MSISDN\tDate\tOPERATOR\tIMSI",
  "separator": "\t",
  "columns":
  [
    {
      "type": "msisdn",
      "format": "(37368|37369|37378)",
      "length": 11
    },
    {
      "type": "date",
      "format": "YYYY-MM-DD hh:mm:ss"
    },
    {
      "type": "string",
      "format": "(ORANGE|NA)"
    },
    {
      "type": "imsi",
      "format": "(25901|25902)"
    }
  ]
}
```
