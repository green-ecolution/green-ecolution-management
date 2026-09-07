---
slug: map
title: Die Karte
part: greenspaces
summary: Bäume und Bewässerungsgruppen auf der Karte lesen, sich orientieren und beides direkt dort erfassen.
routes: ['/map', '/map/tree/new', '/map/treecluster/new']
---

## Die Karte lesen

Die Karte zeigt alle Bäume und Bewässerungsgruppen in deinem Zuständigkeitsbereich an
ihrem tatsächlichen Standort. Beim Herauszoomen fasst sie benachbarte Bäume zu
Gruppensymbolen zusammen; klickst du eines an oder zoomst weiter hinein, löst es sich
in die einzelnen Baumsymbole auf. Jede Gruppe erscheint zusätzlich als eingefärbte
Umrandung um ihre Bäume, ein Klick auf die Fläche öffnet dieselbe Ansicht wie ein Klick
auf das Gruppensymbol.

Ein Klick auf einen Baum führt direkt zu dessen [Detailseite](./trees.md). Ein Klick auf
eine Gruppe öffnet stattdessen eine Seitenansicht mit Adresse, Baumzahl und den
häufigsten Baumarten, den vier Kennzahlen **Bewässerungszustand**, **Bodentemperatur**,
**Letzte Messung** und **Letzte Bewässerung**, sowie einer Vorschau der enthaltenen
Bäume; Bäume mit Sensor sind darin als **Sensor-Baum** markiert, die übrigen mit dem
Hinweis **kein Sensor**. Über **Zum Dashboard** gelangst du von dort zum vollständigen
[Gruppendashboard](./treecluster.md), über das Stiftsymbol daneben direkt in die
Bearbeitung. Oberhalb der Karte filterst du die Anzeige über die Suche nach einer Gruppe
und den Zustand der Bewässerung; wer eine Gruppe anlegen darf, findet dort auch die
Schaltfläche **Gruppe anlegen**.

![Die Kartenübersicht mit Bäumen und Bewässerungsgruppen im Zuständigkeitsbereich, Panele geschlossen](../images/map-overview.png)

## Ebenen und Legende

Über die Schaltfläche mit dem Fragezeichen unten links blendest du die Legende
**Bewässerungsstatus** ein und wieder aus. Sie zeigt dieselben fünf Zustände, die auch
Baum- und Gruppensymbole einfärben:

| Zustand          | Bedeutung                                                      |
| ---------------- | -------------------------------------------------------------- |
| In Ordnung       | Die Bewässerung ist ausreichend, keine Maßnahmen erforderlich. |
| Leicht trocken   | Die Bäume sind leicht trocken und benötigen etwas Wasser.      |
| Sehr trocken     | Die Bäume benötigen dringend Wasser.                           |
| Soeben bewässert | Die Bäume wurden vor kurzem bewässert.                         |
| Unbekannt        | Der Bewässerungsstatus ist unbekannt.                          |

Die ersten drei Zustände bilden die eigentliche Ampel von grün über gelb nach rot, die
beiden übrigen sind neutral eingefärbt. Wie ein Baum und eine Gruppe zu ihrem jeweiligen
Zustand kommen, erklärt das [Gruppendashboard](./treecluster.md#bewasserungsstatus-und-wie-er-zustande-kommt).

![Die eingeblendete Legende Bewässerungsstatus mit ihren fünf Zuständen](../images/map-legend.png)

## Sich auf der Karte bewegen

Die Karte bedienst du vollständig mit der Maus. Halte die linke Maustaste gedrückt und
zieh die Karte an die Stelle, die du sehen willst. Mit dem Mausrad zoomst du hinein und
heraus, und zwar immer auf den Punkt unter dem Mauszeiger: Was du unter dem Cursor hast,
bleibt beim Zoomen an seinem Platz. Dieselben Zoomstufen erreichst du über die
Schaltflächen mit dem Plus und dem Minus in der Kartensteuerung rechts unten und über
einen Doppelklick auf die Karte, der eine Stufe näher heranholt; mit gehaltener
Umschalttaste geht der Doppelklick eine Stufe heraus. Ziehst du mit gehaltener
Umschalttaste ein Rechteck auf, zoomt die Karte genau auf diesen Ausschnitt.

Die Blickrichtung änderst du mit der rechten Maustaste. Halte sie gedrückt und zieh nach
links oder rechts, um die Karte zu drehen, oder nach oben und unten, um sie zu neigen.
Beides zusammen ergibt dieselbe perspektivische Ansicht, die auch die Schaltfläche **3D**
einnimmt. Wenn du die rechte Maustaste nicht verwenden kannst, halte stattdessen die
Strg-Taste und zieh mit der linken.

An einem Tablet oder Telefon verschiebst du die Karte mit einem Finger und zoomst mit der
Aufziehbewegung aus zwei Fingern. Drehst du diese zwei Finger gegeneinander, dreht sich
die Karte mit; ziehst du sie gemeinsam nach oben oder unten, neigt sie sich. Ein doppeltes
Antippen zoomt eine Stufe näher heran, ein Tipp mit zwei Fingern eine Stufe heraus.

Zum Zoomen mit nur einer Hand tippst du doppelt, lässt den Finger beim zweiten Mal aber
liegen: Ziehst du ihn dann nach unten, zoomt die Karte stufenlos hinein, nach oben wieder
heraus.

Auch über die Tastatur lässt sich die Karte führen, sobald sie den Fokus hat. Dafür genügt
ein Klick auf die Karte oder das Weiterspringen mit der Tabulatortaste.

| Taste                              | Wirkung                  |
| ---------------------------------- | ------------------------ |
| Pfeiltasten                        | Karte verschieben        |
| Plus und Minus                     | hinein- und herauszoomen |
| Umschalt + Pfeil links oder rechts | Karte drehen             |
| Umschalt + Pfeil oben oder unten   | Karte neigen             |

> [!NOTE]
> Die Karte bleibt am Rand des Gebiets stehen, das für deine Installation hinterlegt ist,
> und auch der Zoom ist nach beiden Seiten begrenzt. Weiter hinaus geht es nicht, das ist
> kein Fehler.

## Die 3D-Ansicht

Die Schaltfläche **3D** in der Kartensteuerung rechts unten kippt die Karte in eine
perspektivische Ansicht und schaltet auf erneuten Klick wieder in die flache 2D-Ansicht
zurück. Hast du die Karte zusätzlich gedreht, richtet sie die Kompassschaltfläche
darüber mit einem Klick wieder exakt nach Norden aus und hebt die Kippung dabei
gleichzeitig auf.

![Die Karte in der gekippten 3D-Ansicht](../images/map-3d.png)

## Der eigene Standort

Über die Schaltfläche **Eigenen Standort anzeigen** bittet die Anwendung den Browser um
deinen aktuellen Standort und zeigt ihn als Punkt samt Genauigkeitskreis auf der Karte
an; die Kartenansicht folgt deiner Position, bis du selbst an eine andere Stelle
scrollst oder ziehst. Ein erneuter Klick auf dieselbe, jetzt aktive Schaltfläche beendet
die Standortanzeige wieder.

> [!NOTE]
> Unterstützt der Browser keine Standortbestimmung oder liegt die ermittelte Position
> außerhalb des dargestellten Kartenbereichs, meldet dir die Anwendung das direkt und
> schaltet die Standortanzeige wieder aus.

## Bäume und Gruppen direkt auf der Karte erfassen

Einen neuen Baum legst du über die Schaltfläche mit dem Zahnrad-Symbol
(**Kataster-Einstellungen**) oben in der Werkzeugleiste an: Sie öffnet einen Dialog, aus
dem heraus du über **Neuen Baum manuell hinzufügen** zur Erfassung wechselst. Klicke
anschließend auf die Karte, um den Standort des Baums festzulegen; den gesetzten Marker
kannst du danach noch beliebig verschieben. Trage dann Baumnummer, Baumart, Pflanzjahr
und optional eine Bewässerungsgruppe sowie einen Sensor ein und speichere.

![Der Dialog Kataster-Einstellungen mit der Schaltfläche Neuen Baum manuell hinzufügen](../images/map-tree-new.png)

Eine neue Bewässerungsgruppe legst du über die Schaltfläche **Gruppe anlegen** an. Nach
Name, Adresse und Bodenbeschaffenheit klickst du die Bäume, die zur Gruppe gehören
sollen, einzeln auf der Karte an; ein erneuter Klick entfernt einen Baum wieder aus der
Auswahl. Blass dargestellte Bäume gehören einer anderen Organisation und lassen sich
einer neuen Gruppe nicht ohne Weiteres hinzufügen, dazu mehr im
[Kapitel zu den Bewässerungsgruppen](./treecluster.md#baume-zuordnen).

Das Bearbeiten eines Baums oder einer bestehenden Gruppe läuft nach demselben Muster:
Einen Baum rufst du dazu über dessen Detailseite oder direkt über sein Symbol auf der
Karte auf, eine Gruppe über das Stiftsymbol in ihrer Seitenansicht. In beiden Fällen
öffnet sich dieselbe Eingabemaske wie beim Anlegen, bereits mit den vorhandenen Werten
gefüllt.
