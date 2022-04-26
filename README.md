# tcfna
Game engine for "The Campaign for North Africa" as published by Simulations Publications, Inc (SPI).

SPI's [Campaign for North Africa](https://www.spigames.net/campaign_for_north_africa.htm) is a boardgame.

This repository is built to help me run a computer-moderated game.
I'm assuming that all players will download the
[rules, maps, and counters](https://www.spigames.net/PDFv2/CampaignNorthAfrica.pdf)
for themselves.

# Map
Data for generating the map was sourced from Michael Miller's "Hex Database" spreadsheet.
All errors converting to JSON are entirely my fault.

## Player notes for The Campaign for North Africa

The best source for the rules is the SPI site.
The PDF is linked
[here](https://www.spigames.net/PDFv2/CampaignNorthAfrica.pdf)
for convenience.

You will need to print out counters and the map to play.

## Other Resources
The [Campaign for North Africa - Play Group](https://cna-group.proboards.com/) has a forum.

Michael Miller created a very nice
[PDF](https://www.dropbox.com/s/l31ukh5m1kotm5i/CNA-CB-Map-v1.4.pdf?dl=0)
of the map.

SVG Hex logic from
[bencates gist](https://gist.githubusercontent.com/bencates/5b490ed79796cbd35863/raw/044c3b2cdd1b431965d2e285ff0acb936a128393/svgtest.html)
was used to generate the SVG of the board.

Shortest path is worth looking at
[NGraph](https://github.com/anvaka/ngraph.path)
and
[Yet another bidirectional algorithm for shortest paths (Pijls & Post)](https://repub.eur.nl/pub/16100/ei2009-10.pdf).

Color codes from W3Schools
[Color Names](https://www.w3schools.com/colors/colors_names.asp)
and their nice
[RGB Tool](https://www.w3schools.com/colors/colors_rgb.asp).

https://www.jamesshore.com/v2/blog/2022/javascript-colors-and-the-corruption-of-buy-vs-build
