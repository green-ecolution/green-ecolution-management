---
slug: sensor-installation
title: Sensor einbauen und aktivieren
part: sensors
summary: Eine Sensoreinheit am Baum einbauen, vom Bohrloch über die beiden Bodensensoren und das Funkmodul bis zur Aktivierung in der Anwendung.
routes: ['/sensors/new']
---

Eine Sensoreinheit besteht aus zwei Bodensensoren, dem Funkmodul mit Antenne und einer
Bodeneinbaudose, die das Funkmodul dicht und dennoch zugänglich unter der Oberfläche
hält. Eingebaut wird sie am Baum selbst: die Sensoren in zwei Tiefen, das Funkmodul in
der Dose knapp unter dem Niveau der Fläche. Die Aktivierung in der Anwendung gehört
mitten in diesen Ablauf, weil der QR-Code auf dem Funkmodul danach nicht mehr zugänglich
ist.

Mitzunehmen sind neben der Sensoreinheit ein Erdbohrer, ein Leitungssuchgerät und ein
Gerät mit Kamera und Netzverbindung für die Aktivierung, sinnvollerweise mit
[installierter App](./getting-started.md#green-ecolution-als-app-installieren).

![Der Einbau im Schnitt: Bohrloch mit Sensor 2 auf 80 und Sensor 1 auf 40 Zentimetern Tiefe, jeweils im 45-Grad-Winkel zum Baum, darüber die Bodeneinbaudose](../images/sensor-installation.png)

## Bohrloch vorbereiten

Vor dem ersten Spatenstich gehört ein Blick in die Leitungspläne der Fläche, und die
gewählte Einbaustelle wird zusätzlich mit einem Leitungssuchgerät kontrolliert. Erst
danach wird das Loch gebohrt, etwa 80 Zentimeter tief.

> [!WARNING]
> Der Erdbohrer trifft in dieser Tiefe auf Strom-, Wasser- und Telekommunikationsleitungen.
> Die Prüfung mit dem Leitungssuchgerät ersetzt nicht den Leitungsplan und der Plan nicht
> die Prüfung. Beides ist nötig, jedes Mal.

## Sensoren einsetzen

Die beiden Bodensensoren kommen in zwei unterschiedliche Tiefen, jeweils in einem Winkel
von etwa 45 Grad in Richtung des Baums geneigt, damit sie im durchwurzelten Bereich unter
der Krone messen und nicht daneben.

1. Setz den Sensor mit der Nummer 2 auf 80 Zentimetern Tiefe ein, im 45-Grad-Winkel zum Baum geneigt.
2. Füll das Loch auf etwa 40 Zentimeter auf und setz dort den Sensor mit der Nummer 1 ein, in derselben Neigung zum Baum.
3. Füll das Loch weiter auf und setz die Bodeneinbaudose ein, leicht unter der Oberfläche und höchstens etwa 5 Zentimeter tief.
4. Leg die Kabel geordnet in der Bodeneinbaudose ab, ohne sie zu knicken oder zu spannen.

Die Nummerierung der Sensoren ist nicht beliebig: In der Anwendung erscheinen die
Messwerte später je Tiefe, und die Zuordnung ergibt sich aus dieser Reihenfolge. Ein
vertauschter Einbau liefert Messwerte, die vertauscht dargestellt werden, ohne dass es
irgendwo auffällt.

## Funkmodul anschließen

Pack das Funkmodul aus, entferne die gelbe Schutzkappe und schraube die Antenne an.
Verbinde dann das Funkmodul mit den beiden Sensoren und achte dabei auf die Einkerbungen
der Stecker, die nur in einer Stellung zusammenpassen. Der rote Dichtungsring bleibt
dabei am Stecker, und der Schutzverschluss wird fest zugeschraubt. Er hält die
Verbindung auf Jahre trocken.

## Sensor aktivieren

Solange das Funkmodul noch zugänglich ist, wird der Sensor in der Anwendung aktiviert.
Der Aktivierungsassistent unter **Sensor aktivieren** führt in drei Schritten durch das
Verknüpfen einer physischen Sensoreinheit mit einem Baum:

1. **QR-Scan.** Halte den QR-Code auf der Sensoreinheit vor die Kamera. Die Anwendung gleicht die gescannte ID mit dem System ab. Ist der Sensor unbekannt oder die ID ungültig, weist die Anwendung darauf hin und bietet einen erneuten Scan an. Ist der Sensor bereits aktiviert, egal ob **Online** oder **Offline**, lässt er sich hier nicht ein zweites Mal aktivieren; ein bereits verknüpftes Gerät wird stattdessen über seine Detailseite neu zugeordnet, siehe [Sensor deaktivieren und neu verknüpfen](./sensor-readings.md#sensor-deaktivieren-und-neu-verknupfen).
2. **Baum zuordnen.** Die Anwendung ermittelt deinen aktuellen Standort und schlägt Bäume in der Nähe vor, wahlweise auf einer kleinen Karte oder als Liste. Der passende Baum lässt sich auch über die Suche nach Baumnummer oder Baumart finden, falls er nicht unter den Vorschlägen ist. Gespeichert wird am Sensor ausschließlich diese Baumzuordnung, nicht dein GPS-Standort zum Zeitpunkt der Aktivierung; die Position des Sensors auf der Karte ergibt sich später allein aus dem Standort des verknüpften Baums.
3. **Zuordnung prüfen.** Ein letzter Überblick zeigt Sensor-ID, gewählten Baum und deine aktuelle Position zur Kontrolle. Ein Klick auf **Sensor aktivieren** speichert die Verknüpfung; danach lässt sich direkt der nächste Sensor scannen oder zur Übersicht zurückkehren.

![Der Aktivierungsassistent im ersten Schritt: QR-Scan der Sensoreinheit](../images/sensor-wizard-qr.png)

> [!TIP]
> Aktiviere den Sensor, bevor du die Bodeneinbaudose verschließt und die Fläche
> auffüllst. Der QR-Code sitzt auf dem Funkmodul, und ein nachträglicher Scan bedeutet,
> die Dose wieder freizulegen.

## Einbau abschließen

Leg das Funkmodul mit der Antenne nach oben in die Bodeneinbaudose, schließe den grünen
Deckel und füll die Fläche mit höchstens 5 Zentimetern Erde wieder eben auf. Die Antenne
nach oben ist keine Kosmetik: Sie entscheidet darüber, ob die Übertragung das Gateway
überhaupt erreicht.

Die ersten Messwerte können bis zu 12 Stunden brauchen. Danach lohnt der Blick in die
Anwendung: Steht der Sensor auf der Detailseite auf **Online** und trägt er einen
Zeitpunkt unter **Letztes Signal**, ist der Einbau abgeschlossen. Bleibt er auf
**Offline**, hilft das
[Kapitel zur Fehlersuche](./troubleshooting.md) weiter.
