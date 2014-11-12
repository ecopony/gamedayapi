gamedayapi - an mlb gameday api for go
======================================

Just getting going here.

Right now, will just go get a game for the team/date provided and output the game xml to the console.

Prerequisites: go

To build and run
---------

In your go workspace, under a directory github.com/ecopony/

    git clone git@github.com:ecopony/gamedayapi.git
    cd gamedayapi
    go build gameday/mlbgd.go
    ./mlbgd sea 2014-07-22


Usage Restriction
=================

All documents fetched by this API clearly state: Copyright 2011 MLB Advanced Media, L.P. Use of any content on this page
acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt

Which furthermore states: "The accounts, descriptions, data and presentation in the referring page (the “Materials”) are
proprietary content of MLB Advanced Media, L.P (“MLBAM”). Only individual, non-commercial, non-bulk use of the Materials
is permitted and any other use of the Materials is prohibited without prior written authorization from MLBAM. Authorized
users of the Materials are prohibited from using the Materials in any commercial manner other than as
expressly authorized by MLBAM."

Naturally, these terms are passed on to any who use this API. It is your responsibility to abide by them.