package ro

import (
	"github.com/ltraniello/bleve/v2/analysis"
	"github.com/ltraniello/bleve/v2/registry"
)

const StopName = "stop_ro"

// this content was obtained from:
// lucene-4.7.2/analysis/common/src/resources/org/apache/lucene/analysis/
// ` was changed to ' to allow for literal string

var RomanianStopWords = []byte(`# This file was created by Jacques Savoy and is distributed under the BSD license.
# See http://members.unine.ch/jacques.savoy/clef/index.html.
# Also see http://www.opensource.org/licenses/bsd-license.html
acea
aceasta
această
aceea
acei
aceia
acel
acela
acele
acelea
acest
acesta
aceste
acestea
aceşti
aceştia
acolo
acum
ai
aia
aibă
aici
al
ăla
ale
alea
ălea
altceva
altcineva
am
ar
are
aş
aşadar
asemenea
asta
ăsta
astăzi
astea
ăstea
ăştia
asupra
aţi
au
avea
avem
aveţi
azi
bine
bucur
bună
ca
că
căci
când
care
cărei
căror
cărui
cât
câte
câţi
către
câtva
ce
cel
ceva
chiar
cînd
cine
cineva
cît
cîte
cîţi
cîtva
contra
cu
cum
cumva
curând
curînd
da
dă
dacă
dar
datorită
de
deci
deja
deoarece
departe
deşi
din
dinaintea
dintr
dintre
drept
după
ea
ei
el
ele
eram
este
eşti
eu
face
fără
fi
fie
fiecare
fii
fim
fiţi
iar
ieri
îi
îl
îmi
împotriva
în 
înainte
înaintea
încât
încît
încotro
între
întrucât
întrucît
îţi
la
lângă
le
li
lîngă
lor
lui
mă
mâine
mea
mei
mele
mereu
meu
mi
mine
mult
multă
mulţi
ne
nicăieri
nici
nimeni
nişte
noastră
noastre
noi
noştri
nostru
nu
ori
oricând
oricare
oricât
orice
oricînd
oricine
oricît
oricum
oriunde
până
pe
pentru
peste
pînă
poate
pot
prea
prima
primul
prin
printr
sa
să
săi
sale
sau
său
se
şi
sînt
sîntem
sînteţi
spre
sub
sunt
suntem
sunteţi
ta
tăi
tale
tău
te
ţi
ţie
tine
toată
toate
tot
toţi
totuşi
tu
un
una
unde
undeva
unei
unele
uneori
unor
vă
vi
voastră
voastre
voi
voştri
vostru
vouă
vreo
vreun
`)

func TokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	rv := analysis.NewTokenMap()
	err := rv.LoadBytes(RomanianStopWords)
	return rv, err
}

func init() {
	registry.RegisterTokenMap(StopName, TokenMapConstructor)
}
