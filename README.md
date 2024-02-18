# What is eve-online-tools-cli?

`eve-online-tools-cli` is a command-line tool that allows you to retrieve information or do various EVE Online related tasks without being logged into the game.

## Example usage

What you run in the terminal:

```
eve-online-tools api get-market-orders --regionID 10000002
```

What you get:

```
duration,is_buy_order,issued,location_id,min_volume,order_id,price,range,system_id,type_id,volume_remain,volume_total,downloaded_at,region_id
90,false,2024-01-29T15:43:09Z,60003760,1,6703371605,610100000,region,30000142,60755,2,2,2024-02-18T15:07:46Z,10000002
90,false,2023-12-02T08:02:08Z,60003760,1,6652166149,34510,region,30000142,7329,138,145,2024-02-18T15:07:46Z,10000002
...
...
90,true,2024-02-18T10:33:15Z,60003760,1,6718559086,56920,station,30000142,16651,50000,50000,2024-02-18T15:08:27Z,10000002
90,true,2024-02-03T03:47:42Z,60003760,1,6697587567,66000,station,30000142,46934,8,11,2024-02-18T15:08:27Z,10000002
```

### Installation

Go to the [RELEASES](https://github.com/GabrielDCelery/eve-online-tools-cli/releases) page and in the `Assets` section find, download and then extract the `.zip` that is compatible with your system.
