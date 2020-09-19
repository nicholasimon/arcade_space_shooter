package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var ( // MARK: var
	// levels
	levellistmap = make([]int, 200)
	// snow
	snowon      bool
	snowv2map   = make([]rl.Vector2, 300)
	snowtypemap = make([]int, 300)
	// rain
	rainon  bool
	rainmap = make([]rl.Vector2, 500)
	// night
	nighton bool
	// miniboss
	minibossbulletcolor                = rl.Purple
	minibossmodifiers                  minibossmods
	minibosson, createminibosscomplete bool
	minibossbulletmap                  = make([]minibossbullet, 1000)
	minibossbulletcount                int
	// options menu
	optionsmenuselect = 1

	// next level screen
	nextlevelselect = 1
	selectboxborder = rl.Green
	// shop

	shopshipflame                                                                                                                                                                                                                                = rl.NewRectangle(300, 305, 16, 40)
	shopshiplr                                                                                                                                                                                                                                   bool
	shopshipimageufo                                                                                                                                                                                                                             = rl.NewRectangle(200, 305, 91, 91)
	shopshiptarget                                                                                                                                                                                                                               = rl.NewRectangle(200, 320, 112, 75)
	shopshipimage                                                                                                                                                                                                                                = rl.NewRectangle(200, 320, 112, 75)
	currenttarget                                                                                                                                                                                                                                rl.Rectangle
	redshippurchased, greenshippurchased, yellowshippurchased, ufopurchased, target1purchased, target2purchased, target3purchased, target4purchased, greenscannerpurchased, yellowscannerpurchased, redscannerpurchashed, purplescannerpurchased bool
	scannercolor                                                                                                                                                                                                                                 rl.Color
	storeselectfade                                                                                                                                                                                                                              = float32(1.0)
	storeselectfadeon                                                                                                                                                                                                                            bool
	storeselect                                                                                                                                                                                                                                  = 0
	shopon                                                                                                                                                                                                                                       bool
	coinnumber                                                                                                                                                                                                                                   int
	shopscannerradius2                                                                                                                                                                                                                           = float32(50)
	shopscannerradius                                                                                                                                                                                                                            = float32(60)
	shopscannerradiuson                                                                                                                                                                                                                          bool
	// coins
	coinscreated bool
	coinsmap     = make([]coinstruct, 10)
	// credit
	creditscolor  rl.Color
	creditson     bool
	creditsfade   = float32(0.0)
	creditsfadeon bool
	// input
	spacebaroff bool
	// new level
	totalscolor        = rl.Yellow
	resetimerscomplete bool
	startlevelnamefade = float32(1.0)

	startlevelon bool
	stoptimer    bool
	// start menu
	startmenuon     bool
	startmenuselect = 1
	// start game screen
	startscreenship                                                                                                              = rl.NewRectangle(float32(monw/2), float32(monh/2), 112, 75)
	startplayerdirection                                                                                                         int
	startplanet1xchange, startplanet1ychange, startplanet2xchange, startplanet2ychange, startplanet3xchange, startplanet3ychange float32
	startplanetnumber                                                                                                            int
	startgeneratecomplete                                                                                                        bool
	startplanets                                                                                                                 int
	startplanet1, startplanet2, startplanet3                                                                                     rl.Rectangle
	startplanet1v2, startplanet2v2, startplanet3v2                                                                               rl.Vector2
	startplanet1lr, startplanet2lr, startplanet3lr                                                                               bool
	startgamescreenon                                                                                                            bool
	starttitlefade                                                                                                               = float32(0)
	starttitlefadeon                                                                                                             bool
	// player modifiers
	playermodifiers playermod
	// player
	playeroutlinecircleradius   = float32(20)
	playeroutlinecircleradiuson bool
	playerw                     = float32(40)
	playerh                     = float32(40)
	playerhp                    = 100
	playerhptimer               = 3
	playerhppause               bool
	playercollisionrec          rl.Rectangle
	playerx, playery            float32
	playerxy                    rl.Vector2
	playerdirection             = 2
	// options menu
	optionsmenuon bool
	switchclouds  bool
	// bullets
	bulletxycheck   rl.Vector2
	bulletmap       = make([]bullet, 500)
	bulletcount     int
	bulletmodifiers bulletmod
	// level
	tropicalnames    = []string{"tropical", "equatorial", "lush", "luscious", "palm"}
	stonenames       = []string{"stone", "rock", "gravel", "stonework", "masonry", "shale", "pebble", "bedrock", "cobblestone", "slag"}
	sportsnames      = []string{"golf", "athletics", "games", "recreation", "amusement", "exercise", "play", "workout", "training"}
	magentanames     = []string{"rose", "blush", "fuchsia", "roseate", "salmon", "mulberry", "perse", "amaranthine", "lilac", "violet", "plum", "amethyst", "periwinkle"}
	whitenames       = []string{"silver", "silvery", "frosted", "ivory", "alabaster", "blanched", "bleached", "pearly", "achromatic", "neutral", "snowy", "waxen", "milky", "ashen"}
	coldnames        = []string{"icy", "frozen", "chilly", "cold", "frost-bound", "frigid", "glacial", "frosty", "antarctic", "arctic", "shivering", "iced", "polar"}
	fieldnames       = []string{"field", "plain", "meadow", "pasture", "garden", "grassland", "moorland", "terrain", "prairie", "savanna", "steppe"}
	rednames         = []string{"red", "cardinal", "coral", "crimson", "flaming", "glowing", "maroon", "rose", "cerise", "burgundy", "vermilion"}
	bluenames        = []string{"azure", "blue", "teal", "cerulean", "indigo", "beryl", "ultramarine"}
	greennames       = []string{"green", "olive", "jade", "pine", "lime", "malachite", "verdigris", "fir", "viridian", "willow"}
	sadnames         = []string{"bitter", "dismal", "heartbroken", "melancholy", "mournful", "somber", "sorrowful", "dejected", "despondent", "cheerless", "gloomy", "glum", "lugubrious", "morbid", "morose"}
	townnames        = []string{"town", "city", "suburbs", "metropolis", "downtown", "borough", "boondocks", "hamlet", "burg"}
	oddnames         = []string{"odd", "offbeat", "peculiar", "unusual", "oddball", "eccentric", "irregular", "eccentric", "unnatural", "anomalous", "aberrant", "deviant"}
	oasisnames       = []string{"oasis", "fountain", "wellspring", "watering hole", "plateau", "flatland", "tundra", "moor", "heath", "moorland"}
	woodsnames       = []string{"woods", "forest", "grove", "coppice", "woodland", "bosk", "thicket", "scrub", "brushwood", "dingle"}
	psychedelicnames = []string{"psychedelic", "kaleidoscopic", "multicolored", "mottled", "checkered", "dappled", "polychrome", "prismatic", "varicolored", "versicolor"}
	basenames        = []string{"base", "camp", "center", "depot", "garrison", "post", "terminal", "station", "dock", "hangar", "outpost", "habitation", "colony", "encampment"}

	currentlevelname   string
	currentlevelnumber int
	diifculty          = 2
	terraintype        string
	endlevel           bool
	// start level
	startleveltimer  = 1
	startlevelfadeon bool
	startlevelfade   = float32(1.0)
	// end level
	interstitialon bool
	endlevelfade   = float32(0.0)
	endlevelfadeon bool
	// boss level
	bosslinesfade                                                                              = float32(0.4)
	bosslinesfadeon                                                                            bool
	highlightfade1, highlightfade2, highlightfade3, highlightfade4                             float32
	highlightedtileblock1, highlightedtileblock2, highlightedtileblock3, highlightedtileblock4 int
	bosslevelon, bossbackfadecomplete                                                          bool
	bossbackfade                                                                               = float32(0.0)
	// alien code
	aliencodeactivenumber, aliencodeplayercount int
	aliencodemap                                = make([]aliencode, 10)
	aliencodemapplayer                          = make([]aliencode, 3)
	aliencodeactive, aliencodeselected          bool
	// powerups
	poweruptimer1complete, poweruptimer2complete, poweruptimer3complete, poweruptimer4complete, poweruptimer5complete, poweruptimer6complete bool
	poweruptimer1, poweruptimer2, poweruptimer3, poweruptimer4, poweruptimer5, poweruptimer6                                                 int
	powerupcircleradius                                                                                                                      = float32(27)
	powerupmapnumber                                                                                                                         int
	powerupsmap                                                                                                                              = make([]powerup, 2)
	powerupactive                                                                                                                            bool
	powerupnumber                                                                                                                            int
	// colors

	whitefadeon bool
	whitefade1  = float32(0.2)
	whitefade2  = float32(0.3)
	whitefade3  = float32(0.4)
	whitefade4  = float32(0.5)
	whitefade5  = float32(0.6)

	altcloudson bool
	cloudcolor  rl.Color

	randomcolorslice                                                                                                                                                          = make([]rl.Color, 10)
	altforestcolor1, altforestcolor2, altforestcolor3, altforestcolor4, altforestcolor5, altforestcolor6, altforestcolor7, altforestcolor8, altforestcolor9, altforestcolor10 rl.Color

	forestgreen1    = rl.NewColor(46, 176, 180, 255)
	forestgreen2    = rl.NewColor(28, 107, 110, 255)
	forestgreen3    = rl.NewColor(40, 153, 156, 255)
	forestgreen4    = rl.NewColor(35, 136, 139, 255)
	aliencodefade   = float32(1.0)
	aliencodefadeon bool
	powerfadeon     bool
	powerfade       = float32(1.0)
	explosioncolor1 = rl.NewColor(255, 90, 0, 255)
	explosioncolor2 = rl.NewColor(255, 154, 0, 255)
	explosioncolor3 = rl.NewColor(255, 206, 0, 255)
	explosioncolor4 = rl.NewColor(182, 34, 3, 255)
	explosioncolor5 = rl.NewColor(215, 53, 2, 255)
	explosioncolor6 = rl.NewColor(250, 192, 0, 255)
	green1          = rl.NewColor(77, 255, 77, 255)
	bulletcolor1    = rl.NewColor(241, 72, 251, 255)
	bulletfade      = float32(1.0)
	bulletfadeon    bool
	shipflamefade   = float32(1.0)
	// score
	kills, collisions, totalkills, totalcollisions int
	score, totalscore                              int
	scorecolor                                     = rl.Yellow
	xchangescoretext                               int
	// timers
	seconds int
	// enemies
	newenemies bool
	enemytype  string
	newenemyon bool
	enemymap   = make([]enemy, 100)
	// explosions
	explosionsenemycircles                = make([]explosioncircle, 10)
	explosionenemytimer                   = 1
	explosionenemystart, explosionenemyon bool
	explosionenemyx, explosionenemyy      int
	explosionscircles                     = make([]explosioncircle, 10)
	explosiontimer                        = 1
	explosionstart, explosionon           bool
	explosionx, explosiony                int
	// cameras
	cameraclouds rl.Camera2D
	camera4X     rl.Camera2D
	camera2X     rl.Camera2D
	camera15X    rl.Camera2D
	// obstructions
	obstructionrec rl.Rectangle
	collisiontrue  bool
	// fx
	pixelnoiseon                                          bool
	cloud1drift, cloud2drift, cloud3drift                 float32
	cloud1drifton, cloud2drifton, cloud3drifton           bool
	cloud1speed, cloud2speed, cloud3speed                 float32
	cloud1active, cloud2active, cloud3active, startclouds bool
	cloud1v2                                              = rl.NewVector2(-500, -500)
	cloud2v2                                              = rl.NewVector2(-500, -500)
	cloud3v2                                              = rl.NewVector2(-500, -500)
	scanlineson, cloudson, createclouds                   bool
	cloud1lr, cloud2lr, cloud3lr                          bool
	cloudscount                                           int
	// maps
	playerpowerups     = make([]powerup, 5)
	obstructionmap     = make([]bool, blocknumber*2)
	obstructiontypemap = make([]int, blocknumber*2)
	blocktiles         = make([]int, blocknumber*2)
	blockmap           = make([]isoblock, blocknumber*2)
	// images
	sceneryon bool
	// snow
	snow1 = rl.NewRectangle(372, 250, 48, 48)
	snow2 = rl.NewRectangle(420, 250, 48, 48)
	snow3 = rl.NewRectangle(468, 250, 48, 48)
	snow4 = rl.NewRectangle(516, 250, 48, 48)
	// night
	nightshadow = rl.NewRectangle(0, 3296, 800, 800)
	// mini boss
	miniboss1 = rl.NewRectangle(1040, 846, 288, 413)
	miniboss2 = rl.NewRectangle(1337, 879, 291, 176)
	miniboss3 = rl.NewRectangle(1097, 1300, 176, 155)
	// missiles
	missile1 = rl.NewRectangle(192, 342, 20, 35)
	missile2 = rl.NewRectangle(218, 341, 20, 35)
	missile3 = rl.NewRectangle(243, 341, 20, 35)
	missile4 = rl.NewRectangle(270, 339, 19, 40)
	missile5 = rl.NewRectangle(296, 339, 19, 40)
	missile6 = rl.NewRectangle(320, 339, 19, 40)

	missile7  = rl.NewRectangle(484, 363, 24, 46)
	missile8  = rl.NewRectangle(509, 365, 24, 45)
	missile9  = rl.NewRectangle(534, 365, 24, 45)
	missile10 = rl.NewRectangle(562, 369, 22, 38)
	missile11 = rl.NewRectangle(587, 369, 22, 38)
	missile12 = rl.NewRectangle(613, 368, 22, 38)
	// coins
	coin = rl.NewRectangle(41, 378, 16, 16)
	// planets
	planet1 = rl.NewRectangle(2, 1745, 300, 300)
	planet2 = rl.NewRectangle(321, 1747, 300, 300)
	planet3 = rl.NewRectangle(670, 1747, 300, 300)
	planet4 = rl.NewRectangle(1018, 1746, 300, 300)
	planet5 = rl.NewRectangle(1337, 1745, 300, 300)
	planet6 = rl.NewRectangle(1653, 1746, 300, 300)
	planet7 = rl.NewRectangle(1973, 1747, 300, 300)
	// bullets
	bullet1 = rl.NewRectangle(41, 348, 20, 20)
	bullet2 = rl.NewRectangle(362, 321, 20, 20)
	bullet3 = rl.NewRectangle(362, 351, 20, 20)
	bullet4 = rl.NewRectangle(473, 322, 20, 20)
	// bosses
	boss1 = rl.NewRectangle(1696, 669, 348, 478)
	// code
	code1  = rl.NewRectangle(7, 483, 60, 98)
	code2  = rl.NewRectangle(74, 482, 60, 98)
	code3  = rl.NewRectangle(141, 481, 60, 98)
	code4  = rl.NewRectangle(206, 482, 60, 98)
	code5  = rl.NewRectangle(269, 483, 60, 98)
	code6  = rl.NewRectangle(335, 482, 60, 98)
	code7  = rl.NewRectangle(404, 481, 60, 98)
	code8  = rl.NewRectangle(472, 483, 60, 98)
	code9  = rl.NewRectangle(544, 483, 60, 98)
	code10 = rl.NewRectangle(617, 481, 60, 98)
	// enemies
	enemy1            = rl.NewRectangle(0, 136, 103, 84)
	enemy2            = rl.NewRectangle(113, 137, 93, 84)
	enemy3            = rl.NewRectangle(219, 134, 104, 84)
	enemy4            = rl.NewRectangle(339, 138, 97, 84)
	enemy5            = rl.NewRectangle(934, 706, 102, 83)
	enemy6            = rl.NewRectangle(1071, 706, 108, 83)
	enemy7            = rl.NewRectangle(1211, 701, 100, 96)
	enemy8            = rl.NewRectangle(1331, 708, 116, 84)
	enemy9            = rl.NewRectangle(1467, 708, 103, 77)
	enemyobstruction1 = rl.NewRectangle(0, 245, 82, 84)
	enemyobstruction2 = rl.NewRectangle(85, 245, 82, 84)
	enemyobstruction3 = rl.NewRectangle(171, 244, 82, 84)
	enemyobstruction4 = rl.NewRectangle(257, 243, 82, 84)
	// player
	playership     rl.Rectangle
	lasermarker1   = rl.NewRectangle(0, 0, 37, 38)
	lasermarker2   = rl.NewRectangle(39, 0, 37, 38)
	lasermarker3   = rl.NewRectangle(77, 0, 48, 46)
	lasermarker4   = rl.NewRectangle(126, 0, 48, 46)
	lasermarker5   = rl.NewRectangle(179, 0, 59, 40)
	lasermarker590 = rl.NewRectangle(157, 340, 19, 63)
	ship1          = rl.NewRectangle(0, 47, 112, 75)
	ship2          = rl.NewRectangle(115, 47, 112, 75)
	ship3          = rl.NewRectangle(229, 48, 112, 75)
	ship4          = rl.NewRectangle(344, 51, 112, 75)
	ufo            = rl.NewRectangle(464, 50, 91, 91)
	shipflame      = rl.NewRectangle(0, 356, 16, 40)
	// powerups
	powerup1  = rl.NewRectangle(0, 411, 57, 57)
	powerup2  = rl.NewRectangle(58, 411, 57, 57)
	powerup3  = rl.NewRectangle(115, 411, 57, 57)
	powerup4  = rl.NewRectangle(175, 411, 57, 57)
	powerup5  = rl.NewRectangle(237, 411, 57, 57)
	powerup6  = rl.NewRectangle(297, 413, 57, 57)
	powerup7  = rl.NewRectangle(363, 414, 57, 57)
	powerup8  = rl.NewRectangle(428, 413, 57, 57)
	powerup9  = rl.NewRectangle(490, 414, 57, 57)
	powerup10 = rl.NewRectangle(551, 415, 57, 57)
	// clouds
	cloud1 = rl.NewRectangle(0, 876, 456, 148)
	cloud2 = rl.NewRectangle(459, 878, 426, 146)
	cloud3 = rl.NewRectangle(0, 694, 398, 170)
	// terrain
	// palm island
	palmisland1 = rl.NewRectangle(2368, 1118, 132, 126)
	palmisland2 = rl.NewRectangle(2523, 1106, 132, 138)
	palmisland3 = rl.NewRectangle(2674, 1090, 132, 155)
	palmisland4 = rl.NewRectangle(2825, 1086, 132, 156)
	palmisland5 = rl.NewRectangle(2963, 1077, 132, 165)
	palmisland6 = rl.NewRectangle(3110, 1112, 132, 130)
	// block forest
	blockforest1 = rl.NewRectangle(2361, 897, 132, 152)
	blockforest2 = rl.NewRectangle(2508, 897, 132, 148)
	blockforest3 = rl.NewRectangle(2649, 907, 132, 142)
	blockforest4 = rl.NewRectangle(2800, 897, 132, 151)
	blockforest5 = rl.NewRectangle(2943, 892, 132, 153)
	// odd red
	plainred = rl.NewRectangle(2363, 561, 132, 79)
	oddred1  = rl.NewRectangle(2509, 565, 132, 70)
	oddred2  = rl.NewRectangle(2643, 636, 132, 83)
	oddred3  = rl.NewRectangle(2793, 567, 132, 70)
	oddred4  = rl.NewRectangle(2659, 659, 137, 160)
	oddred5  = rl.NewRectangle(2817, 663, 137, 160)
	// golf
	plaingreen1 = rl.NewRectangle(2522, 478, 132, 72)
	plaingreen2 = rl.NewRectangle(3083, 974, 132, 72)
	plaingreen3 = rl.NewRectangle(3226, 976, 132, 72)
	plaingreen4 = rl.NewRectangle(3371, 976, 132, 72)
	golf1       = rl.NewRectangle(2373, 479, 132, 72)
	golf2       = rl.NewRectangle(2664, 476, 132, 71)
	golf3       = rl.NewRectangle(2359, 674, 132, 131)
	golf4       = rl.NewRectangle(2504, 676, 132, 131)
	// brown stone
	brownstone1 = rl.NewRectangle(2872, 13, 132, 75)
	brownstone2 = rl.NewRectangle(3015, 10, 132, 78)
	brownstone3 = rl.NewRectangle(3156, 14, 132, 76)
	brownstone4 = rl.NewRectangle(3295, 13, 132, 76)
	brownstone5 = rl.NewRectangle(3435, 17, 132, 76)
	brownstone6 = rl.NewRectangle(3571, 19, 132, 76)
	brownstone7 = rl.NewRectangle(3717, 23, 132, 72)
	// trees
	pinetree1  = rl.NewRectangle(576, 21, 54, 161)
	pinetree2  = rl.NewRectangle(639, 13, 55, 174)
	pinetree3  = rl.NewRectangle(712, 12, 55, 174)
	pinetree4  = rl.NewRectangle(793, 14, 57, 175)
	pinetree5  = rl.NewRectangle(870, 16, 56, 135)
	pinetree6  = rl.NewRectangle(941, 39, 58, 92)
	pinetree7  = rl.NewRectangle(1019, 38, 56, 100)
	pinetree8  = rl.NewRectangle(1091, 14, 59, 141)
	pinetree9  = rl.NewRectangle(1163, 15, 57, 144)
	pinetree10 = rl.NewRectangle(1241, 24, 75, 120)
	// mars
	marsflat   = rl.NewRectangle(2059, 6, 132, 66)
	marsrock1  = rl.NewRectangle(2202, 5, 88, 67)
	marsrock2  = rl.NewRectangle(2318, 2, 92, 76)
	marsrock3  = rl.NewRectangle(2423, 3, 92, 81)
	marsrock4  = rl.NewRectangle(2532, 20, 85, 66)
	marsrock5  = rl.NewRectangle(2644, 6, 84, 83)
	marsrock6  = rl.NewRectangle(2754, 8, 84, 83)
	turret1    = rl.NewRectangle(2086, 106, 100, 99)
	turret2    = rl.NewRectangle(2215, 108, 83, 104)
	radar1     = rl.NewRectangle(2328, 122, 78, 95)
	radar2     = rl.NewRectangle(2435, 133, 77, 82)
	radar3     = rl.NewRectangle(2846, 144, 69, 80)
	radar4     = rl.NewRectangle(2577, 342, 77, 67)
	radar5     = rl.NewRectangle(3724, 148, 66, 83)
	structure1 = rl.NewRectangle(2536, 102, 115, 141)
	structure2 = rl.NewRectangle(2678, 116, 132, 188)
	structure3 = rl.NewRectangle(3224, 119, 127, 147)
	structure4 = rl.NewRectangle(3393, 124, 127, 147)
	structure5 = rl.NewRectangle(2946, 134, 125, 101)
	circular1  = rl.NewRectangle(3101, 119, 80, 135)
	circular2  = rl.NewRectangle(3565, 129, 97, 139)
	circular3  = rl.NewRectangle(3828, 127, 116, 138)
	circular4  = rl.NewRectangle(2890, 297, 96, 120)
	hangar1    = rl.NewRectangle(3017, 285, 148, 148)
	hangar2    = rl.NewRectangle(3196, 284, 148, 148)
	hangar3    = rl.NewRectangle(2366, 278, 153, 135)
	marsship1  = rl.NewRectangle(2119, 286, 176, 131)
	marsship2  = rl.NewRectangle(2135, 461, 195, 126)
	// rocky
	rocky1 = rl.NewRectangle(1381, 0, 132, 98)
	rocky2 = rl.NewRectangle(1515, 0, 132, 102)
	rocky3 = rl.NewRectangle(1381, 99, 132, 99)
	rocky4 = rl.NewRectangle(1515, 102, 132, 101)
	rocky5 = rl.NewRectangle(1514, 204, 132, 99)
	rocky6 = rl.NewRectangle(1381, 298, 132, 125)
	rocky7 = rl.NewRectangle(1514, 303, 132, 99)
	rocky8 = rl.NewRectangle(1515, 402, 132, 99)
	// city
	building1 = rl.NewRectangle(935, 171, 132, 139)
	building2 = rl.NewRectangle(1067, 175, 132, 139)
	building3 = rl.NewRectangle(1204, 177, 132, 139)
	building4 = rl.NewRectangle(929, 349, 132, 139)
	building5 = rl.NewRectangle(1062, 352, 132, 139)
	building6 = rl.NewRectangle(1197, 351, 132, 139)
	building7 = rl.NewRectangle(936, 530, 132, 139)
	building8 = rl.NewRectangle(1068, 533, 132, 139)
	building9 = rl.NewRectangle(1199, 534, 132, 139)
	// urban
	concrete         = rl.NewRectangle(1486, 514, 132, 102)
	concretetree     = rl.NewRectangle(1623, 499, 132, 114)
	concretefountain = rl.NewRectangle(1766, 509, 132, 102)
	concretepark     = rl.NewRectangle(1906, 506, 132, 102)
	// woods
	grass   = rl.NewRectangle(1381, 425, 132, 82)
	trees1  = rl.NewRectangle(1648, 0, 132, 118)
	trees2  = rl.NewRectangle(1781, 0, 132, 127)
	trees3  = rl.NewRectangle(1914, 0, 132, 118)
	trees4  = rl.NewRectangle(1647, 118, 132, 123)
	trees5  = rl.NewRectangle(1780, 127, 132, 121)
	trees6  = rl.NewRectangle(1914, 118, 132, 117)
	trees7  = rl.NewRectangle(1648, 241, 132, 112)
	trees8  = rl.NewRectangle(1781, 248, 132, 124)
	trees9  = rl.NewRectangle(1648, 352, 132, 113)
	trees10 = rl.NewRectangle(1781, 371, 132, 122)
	// isometric block grid
	drawblock, nextblock              int
	screenblocknumber                 int
	gridon                            bool
	vertcount, horizcount, gridlayout int
	blockw                            = 132
	blockh                            = 66
	blocknumber                       int
	// core
	fps            = 30
	pauseon        bool
	mouseblock     int
	framecount     int
	debugon        bool
	monh, monw     int
	monh32, monw32 int32
	mousepos       rl.Vector2
	imgs           rl.Texture2D
	camera         rl.Camera2D
)

type minibossmods struct {
	hp, firetype, minibosstype int
	speed                      float32
	v2, v22, v23               rl.Vector2
}
type aliencode struct {
	img                           rl.Rectangle
	created, collected, displayed bool
	v2                            rl.Vector2
}

type powerup struct {
	v2        rl.Vector2
	img       rl.Rectangle
	name      string
	direction int
	side      bool
	active    bool

	angleon, angle2on                   bool
	anglemultiplier, angle              int
	number, spread                      float32
	left, right, down                   bool
	leftnumber, rightnumber, downnumber int

	randommod1, randommod2, randommod3, randommod4, randommod5, randommod6 bool
}
type minibossbullet struct {
	v2            rl.Vector2
	hp, direction int
	active        bool
	img           rl.Rectangle
}
type bullet struct {
	side              int
	damage            int
	xy                rl.Vector2
	active            bool
	left, right, down bool
	xchange           float32
}
type bulletmod struct {
	angleon, angle2on                   bool
	anglemultiplier, angle              int
	number, spread                      float32
	left, right, down                   bool
	leftnumber, rightnumber, downnumber int
	img                                 int
}
type playermod struct {
	shieldon                                              bool
	shieldhp                                              int
	movementdistortionamount, movementdistortionleftright int
	movementdistortionon                                  bool
}

type explosioncircle struct {
	x, y         int
	radius, fade float32
	color        rl.Color
}

type isoblock struct {
	xy, xy2, xy3, xy4, topleft rl.Vector2
}
type enemy struct {
	xy        rl.Vector2
	hp, drift int
	score     int
	img       rl.Rectangle
	active    bool
}

type coinstruct struct {
	xy                rl.Vector2
	active, collected bool
	time              int
	rec               rl.Rectangle
}

/*


 */

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "isometric")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(fps)
	// rl.HideCursor()
	// 	rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {
		mousepos = rl.GetMousePosition()

		checkblock := blockmap[nextblock]
		camera.Target.Y = checkblock.topleft.Y + 33

		framecount++
		if stoptimer == false {
			if framecount%30 == 0 {
				seconds++
				if seconds == 60 { // level time
					endlevel = true
					resetimerscomplete = false
					spacebaroff = true
					stoptimer = true
					if coinnumber >= 5 {
						nextlevelselect = 0
					} else if coinnumber < 5 {
						nextlevelselect = 1
					}
					currentlevelnumber++
				}
			}
		}
		if pauseon == false {
			if framecount%2 == 0 {
				if nextblock > horizcount*4 {
					nextblock -= 34
					camera.Target.Y -= 66
				}
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		// draw visible blocks
		drawblock = nextblock
		for a := 0; a < screenblocknumber; a++ {
			checkblock := blockmap[drawblock]
			checktiles := blocktiles[drawblock]
			checkobstruction := obstructionmap[drawblock]
			//	rl.DrawTextureRec(imgs, grassblock1, checkblock.topleft, rl.White)

			switch terraintype {
			case "magenta":
				switch checktiles {
				case 1:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, whitefade1))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, whitefade1))
				case 2:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, whitefade2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, whitefade2))
				case 3:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, whitefade3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, whitefade3))
				case 4:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, whitefade4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, whitefade4))
				case 5:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, whitefade5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, whitefade5))
				case 6:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, 0.7))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, 0.7))
				case 7:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, 0.8))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, 0.8))
				case 8:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Magenta, 0.9))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Magenta, 0.9))
				}
			case "white":
				switch checktiles {
				case 1:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, whitefade1))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, whitefade1))
				case 2:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, whitefade2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, whitefade2))
				case 3:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, whitefade3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, whitefade3))
				case 4:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, whitefade4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, whitefade4))
				case 5:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, whitefade5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, whitefade5))
				case 6:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, 0.7))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, 0.7))
				case 7:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, 0.8))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, 0.8))
				case 8:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, 0.9))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, 0.9))
				}
			case "oddred":
				switch checktiles {
				case 1:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plainred, v2, rl.Maroon)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, oddred1, v2, rl.Maroon)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plainred, v2, rl.Maroon)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, oddred3, v2, rl.Maroon)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-90)
					rl.DrawTextureRec(imgs, oddred4, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-90)
					rl.DrawTextureRec(imgs, oddred5, v2, rl.White)
				}
			case "golf":
				switch checktiles {
				case 1:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen1, v2, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen2, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen3, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen4, v2, rl.White)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y+1)
					rl.DrawTextureRec(imgs, golf1, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, golf2, v2, rl.White)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-60)
					rl.DrawTextureRec(imgs, golf3, v2, rl.White)
				case 8:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-60)
					rl.DrawTextureRec(imgs, golf4, v2, rl.White)
				}
			case "blockforest":
				switch checktiles {
				case 1:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen1, v2, rl.Magenta)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen2, v2, rl.Magenta)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen3, v2, rl.Magenta)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen4, v2, rl.Magenta)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-81)
					rl.DrawTextureRec(imgs, blockforest1, v2, rl.Magenta)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-80)
					rl.DrawTextureRec(imgs, blockforest2, v2, rl.Magenta)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-67)
					rl.DrawTextureRec(imgs, blockforest3, v2, rl.Magenta)
				case 8:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-79)
					rl.DrawTextureRec(imgs, blockforest4, v2, rl.Magenta)
				case 9:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-80)
					rl.DrawTextureRec(imgs, blockforest5, v2, rl.Magenta)
				}
			case "palmisland":
				switch checktiles {
				case 1:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen1, v2, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen2, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen3, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, plaingreen4, v2, rl.White)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-54)
					rl.DrawTextureRec(imgs, palmisland1, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-66)
					rl.DrawTextureRec(imgs, palmisland2, v2, rl.White)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-82)
					rl.DrawTextureRec(imgs, palmisland3, v2, rl.White)
				case 8:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-84)
					rl.DrawTextureRec(imgs, palmisland4, v2, rl.White)
				case 9:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-94)
					rl.DrawTextureRec(imgs, palmisland5, v2, rl.White)
				case 10:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-58)
					rl.DrawTextureRec(imgs, palmisland6, v2, rl.White)
				}
			case "brownstone":
				switch checktiles {
				case 1:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone1, v2, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone2, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone3, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone4, v2, rl.White)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone5, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone6, v2, rl.White)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, brownstone7, v2, rl.White)
				}
			case "mars4":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 2:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 3:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 4:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 5:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 6:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 7:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 8:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 9:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 10:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 11:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 12:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 13:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 14:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 15:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 16:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 17:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 18:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 19:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 20:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 21:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 22:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 23:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 24:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 25:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 26:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 27:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 28:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 29:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 30:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 31:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+28, checkblock.topleft.Y-85)
					rl.DrawTextureRec(imgs, circular1, v2, rl.White)
				case 32:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+10, checkblock.topleft.Y-90)
					rl.DrawTextureRec(imgs, circular2, v2, rl.White)
				case 33:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+10, checkblock.topleft.Y-90)
					rl.DrawTextureRec(imgs, circular3, v2, rl.White)
				case 34:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+18, checkblock.topleft.Y-60)
					rl.DrawTextureRec(imgs, circular4, v2, rl.White)
				case 35:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+15, checkblock.topleft.Y-45)
					rl.DrawTextureRec(imgs, turret1, v2, rl.White)
				case 36:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+25, checkblock.topleft.Y-50)
					rl.DrawTextureRec(imgs, turret2, v2, rl.White)
				case 37:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-30)
					rl.DrawTextureRec(imgs, marsrock6, v2, rl.White)
				case 38:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-10)
					rl.DrawTextureRec(imgs, marsrock4, v2, rl.White)
				}
			case "mars3":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 2:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 3:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 4:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 5:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 6:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 7:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 8:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 9:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 10:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 11:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 12:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 13:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 14:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 15:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 16:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 17:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 18:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 19:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 20:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 21:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 22:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 23:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 24:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 25:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 26:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 27:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 28:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 29:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 30:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 31:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+6, checkblock.topleft.Y-78)
					rl.DrawTextureRec(imgs, structure1, v2, rl.White)
				case 32:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-148)
					rl.DrawTextureRec(imgs, structure5, v2, rl.White)
				case 33:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-82)
					rl.DrawTextureRec(imgs, structure3, v2, rl.White)
				case 34:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+2, checkblock.topleft.Y-82)
					rl.DrawTextureRec(imgs, structure4, v2, rl.White)
				case 35:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-40)
					rl.DrawTextureRec(imgs, structure5, v2, rl.White)
				case 36:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+22, checkblock.topleft.Y-24)
					rl.DrawTextureRec(imgs, marsrock2, v2, rl.White)
				case 37:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-30)
					rl.DrawTextureRec(imgs, marsrock6, v2, rl.White)
				case 38:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-10)
					rl.DrawTextureRec(imgs, marsrock4, v2, rl.White)
				}
			case "mars2":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 2:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 3:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 4:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 5:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 6:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 7:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 8:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 9:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 10:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 11:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 12:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 13:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 14:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 15:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 16:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 17:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 18:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 19:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 20:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 21:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 22:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 23:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 24:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 25:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 26:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 27:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 28:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 29:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 30:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 31:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+22, checkblock.topleft.Y-45)
					rl.DrawTextureRec(imgs, radar1, v2, rl.White)
				case 32:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+25, checkblock.topleft.Y-25)
					rl.DrawTextureRec(imgs, radar2, v2, rl.White)
				case 33:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+28, checkblock.topleft.Y-25)
					rl.DrawTextureRec(imgs, radar3, v2, rl.White)
				case 34:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+26, checkblock.topleft.Y-15)
					rl.DrawTextureRec(imgs, radar4, v2, rl.White)
				case 35:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+28, checkblock.topleft.Y-30)
					rl.DrawTextureRec(imgs, radar5, v2, rl.White)
				case 36:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+22, checkblock.topleft.Y-24)
					rl.DrawTextureRec(imgs, marsrock2, v2, rl.White)
				case 37:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-30)
					rl.DrawTextureRec(imgs, marsrock6, v2, rl.White)
				case 38:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-10)
					rl.DrawTextureRec(imgs, marsrock4, v2, rl.White)
				}
			case "mars1":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 2:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 3:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 4:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 5:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 6:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 7:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 8:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 9:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 10:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 11:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 12:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 13:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 14:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
				case 15:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+18, checkblock.topleft.Y-8)
					rl.DrawTextureRec(imgs, marsrock1, v2, rl.White)
				case 16:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+22, checkblock.topleft.Y-24)
					rl.DrawTextureRec(imgs, marsrock2, v2, rl.White)
				case 17:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+18, checkblock.topleft.Y-22)
					rl.DrawTextureRec(imgs, marsrock3, v2, rl.White)
				case 18:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-10)
					rl.DrawTextureRec(imgs, marsrock4, v2, rl.White)
				case 19:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-30)
					rl.DrawTextureRec(imgs, marsrock5, v2, rl.White)
				case 20:
					rl.DrawTextureRec(imgs, marsflat, checkblock.topleft, rl.White)
					v2 := rl.NewVector2(checkblock.topleft.X+24, checkblock.topleft.Y-30)
					rl.DrawTextureRec(imgs, marsrock6, v2, rl.White)
				}
			case "forest":
				switch checktiles {
				case 1:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-25, checkblock.xy2.Y-123)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.2))
					rl.DrawTextureRec(imgs, pinetree1, pinetreev2, rl.White)
				case 2:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-130)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.3))
					rl.DrawTextureRec(imgs, pinetree2, pinetreev2, rl.White)
				case 3:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-130)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.4))
					rl.DrawTextureRec(imgs, pinetree3, pinetreev2, rl.White)
				case 4:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-130)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.5))
					rl.DrawTextureRec(imgs, pinetree4, pinetreev2, rl.White)
				case 5:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-100)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.2))
					rl.DrawTextureRec(imgs, pinetree5, pinetreev2, rl.White)
				case 6:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-29, checkblock.xy2.Y-50)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.3))
					rl.DrawTextureRec(imgs, pinetree6, pinetreev2, rl.White)
				case 7:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-29, checkblock.xy2.Y-52)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.4))
					rl.DrawTextureRec(imgs, pinetree7, pinetreev2, rl.White)
				case 8:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-23, checkblock.xy2.Y-98)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.5))
					rl.DrawTextureRec(imgs, pinetree8, pinetreev2, rl.White)
				case 9:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-23, checkblock.xy2.Y-98)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.4))
					rl.DrawTextureRec(imgs, pinetree9, pinetreev2, rl.White)
				case 10:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-40, checkblock.xy2.Y-78)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.2))
					rl.DrawTextureRec(imgs, pinetree10, pinetreev2, rl.White)
				case 11:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.2))
				case 12:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.3))
				case 13:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.4))
				case 14:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.5))
				case 15:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.2))
				case 16:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.3))
				case 17:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.4))
				case 18:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.5))
				case 19:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.4))
				case 20:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.2))
				}
			case "altforest":
				switch checktiles {
				case 1:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-25, checkblock.xy2.Y-123)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.2))
					rl.DrawTextureRec(imgs, pinetree1, pinetreev2, altforestcolor1)
				case 2:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-130)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.3))
					rl.DrawTextureRec(imgs, pinetree2, pinetreev2, altforestcolor2)
				case 3:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-130)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.4))
					rl.DrawTextureRec(imgs, pinetree3, pinetreev2, altforestcolor3)
				case 4:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-130)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.5))
					rl.DrawTextureRec(imgs, pinetree4, pinetreev2, altforestcolor4)
				case 5:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-28, checkblock.xy2.Y-100)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.2))
					rl.DrawTextureRec(imgs, pinetree5, pinetreev2, altforestcolor5)
				case 6:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-29, checkblock.xy2.Y-50)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.3))
					rl.DrawTextureRec(imgs, pinetree6, pinetreev2, altforestcolor6)
				case 7:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-29, checkblock.xy2.Y-52)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.4))
					rl.DrawTextureRec(imgs, pinetree7, pinetreev2, altforestcolor7)
				case 8:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-23, checkblock.xy2.Y-98)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.5))
					rl.DrawTextureRec(imgs, pinetree8, pinetreev2, altforestcolor8)
				case 9:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-23, checkblock.xy2.Y-98)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.4))
					rl.DrawTextureRec(imgs, pinetree9, pinetreev2, altforestcolor9)
				case 10:
					pinetreev2 := rl.NewVector2(checkblock.xy2.X-40, checkblock.xy2.Y-78)
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.2))
					rl.DrawTextureRec(imgs, pinetree10, pinetreev2, altforestcolor10)
				case 11:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.2))
				case 12:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.3))
				case 13:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.4))
				case 14:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen1, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen1, 0.5))
				case 15:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.2))
				case 16:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen2, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen2, 0.3))
				case 17:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.4))
				case 18:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen3, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen3, 0.5))
				case 19:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.4))
				case 20:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(forestgreen4, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(forestgreen4, 0.2))
				}
			case "rocky":
				switch checktiles {
				case 1:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-28)
					rl.DrawTextureRec(imgs, rocky6, v2, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-4)
					rl.DrawTextureRec(imgs, rocky2, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, rocky3, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-4)
					rl.DrawTextureRec(imgs, rocky4, v2, rl.White)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-1)
					rl.DrawTextureRec(imgs, rocky5, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-1)
					rl.DrawTextureRec(imgs, rocky1, v2, rl.White)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-2)
					rl.DrawTextureRec(imgs, rocky7, v2, rl.White)
				case 8:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, rocky8, v2, rl.White)
				}
			case "city":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, concrete, checkblock.topleft, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building1, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building2, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building3, v2, rl.White)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building4, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building5, v2, rl.White)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building6, v2, rl.White)
				case 8:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building7, v2, rl.White)
				case 9:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building8, v2, rl.White)
				case 10:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-36)
					rl.DrawTextureRec(imgs, building9, v2, rl.White)
				case 11:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-12)
					rl.DrawTextureRec(imgs, concretetree, v2, rl.White)
				case 12:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-12)
					rl.DrawTextureRec(imgs, concretetree, v2, rl.White)
				}
			case "urban":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, concrete, checkblock.topleft, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-20)
					rl.DrawTextureRec(imgs, concretetree, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y)
					rl.DrawTextureRec(imgs, concretepark, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y+6)
					rl.DrawTextureRec(imgs, concretefountain, v2, rl.White)
				}
			case "meadow":
				switch checktiles {
				case 1:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Green, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Green, 0.2))
				case 2:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Green, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Green, 0.3))
				case 3:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Green, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Green, 0.4))
				case 4:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Green, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Green, 0.5))
				case 5:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Lime, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Lime, 0.2))
				case 6:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Lime, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Lime, 0.3))
				case 7:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Lime, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Lime, 0.4))
				case 8:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Lime, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Lime, 0.5))
				}
			case "woods":
				switch checktiles {
				case 1:
					rl.DrawTextureRec(imgs, grass, checkblock.topleft, rl.White)
				case 2:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-20)
					rl.DrawTextureRec(imgs, trees1, v2, rl.White)
				case 3:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-29)
					rl.DrawTextureRec(imgs, trees2, v2, rl.White)
				case 4:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-20)
					rl.DrawTextureRec(imgs, trees3, v2, rl.White)
				case 5:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-25)
					rl.DrawTextureRec(imgs, trees4, v2, rl.White)
				case 6:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-23)
					rl.DrawTextureRec(imgs, trees5, v2, rl.White)
				case 7:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-18)
					rl.DrawTextureRec(imgs, trees6, v2, rl.White)
				case 8:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-13)
					rl.DrawTextureRec(imgs, trees7, v2, rl.White)
				case 9:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-26)
					rl.DrawTextureRec(imgs, trees8, v2, rl.White)
				case 10:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-15)
					rl.DrawTextureRec(imgs, trees9, v2, rl.White)
				case 11:
					v2 := rl.NewVector2(checkblock.topleft.X, checkblock.topleft.Y-24)
					rl.DrawTextureRec(imgs, trees10, v2, rl.White)
				}
			case "red":
				switch checktiles {
				case 1:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Red, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Red, 0.2))
				case 2:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Red, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Red, 0.3))
				case 3:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Red, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Red, 0.4))
				case 4:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Red, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Red, 0.5))
				case 5:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Maroon, 0.2))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Maroon, 0.2))
				case 6:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Maroon, 0.3))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Maroon, 0.3))
				case 7:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Maroon, 0.4))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Maroon, 0.4))
				case 8:
					rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Maroon, 0.5))
					rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Maroon, 0.5))
				}
			}

			if gridon {
				rl.DrawLineV(checkblock.xy, checkblock.xy2, rl.Black)
				rl.DrawLineV(checkblock.xy2, checkblock.xy3, rl.Black)
				rl.DrawLineV(checkblock.xy, checkblock.xy4, rl.Black)
				rl.DrawLineV(checkblock.xy4, checkblock.xy3, rl.Black)
			}
			// player ground object collisions
			if pauseon == false {
				if checkobstruction {
					obstructionrec = rl.NewRectangle(checkblock.xy2.X-20, checkblock.xy2.Y+8, 82, 84)
					obsv2 := rl.NewVector2(checkblock.xy2.X-20, checkblock.xy2.Y-50)
					obsshadowv2 := rl.NewVector2(checkblock.xy2.X-15, checkblock.xy2.Y-55)
					switch obstructiontypemap[drawblock] {
					case 1:
						rl.DrawTextureRec(imgs, enemyobstruction1, obsshadowv2, rl.Black)
						rl.DrawTextureRec(imgs, enemyobstruction1, obsv2, rl.White)
					case 2:
						rl.DrawTextureRec(imgs, enemyobstruction2, obsshadowv2, rl.Black)
						rl.DrawTextureRec(imgs, enemyobstruction2, obsv2, rl.White)
					case 3:
						rl.DrawTextureRec(imgs, enemyobstruction3, obsshadowv2, rl.Black)
						rl.DrawTextureRec(imgs, enemyobstruction3, obsv2, rl.White)
					case 4:
						rl.DrawTextureRec(imgs, enemyobstruction4, obsshadowv2, rl.Black)
						rl.DrawTextureRec(imgs, enemyobstruction4, obsv2, rl.White)
					}
				}
				playercollisionrec = rl.NewRectangle(playerx+66, playery+camera.Target.Y, 112, 75)
				if rl.CheckCollisionRecs(playercollisionrec, obstructionrec) {
					collisiontrue = true
					if playerhppause == false {
						if playership == ufo || playership == ship2 {
							playerhp -= 4
						} else {
							playerhp -= 5
						}
						playerhppause = true
						explosionstart = false
						explosion()
						collisions++
					}
				}
			}
			drawblock++
		}

		rl.EndMode2D() // MARK: draw no camera

		// update events
		update()

		rl.EndDrawing()
	}
	rl.CloseWindow()

}
func randomcolors() {

	color1 := rl.Red
	color2 := rl.Orange
	color3 := rl.Yellow
	color4 := rl.Lime
	color5 := rl.Blue
	color6 := rl.Magenta
	color7 := rl.Pink
	color8 := rl.Purple
	color9 := rl.Maroon
	color10 := rl.SkyBlue

	count := 0

	for {
		choose := rInt(1, 11)
		switch choose {
		case 1:
			randomcolorslice[count] = color1
		case 2:
			randomcolorslice[count] = color2
		case 3:
			randomcolorslice[count] = color3
		case 4:
			randomcolorslice[count] = color4
		case 5:
			randomcolorslice[count] = color5
		case 6:
			randomcolorslice[count] = color6
		case 7:
			randomcolorslice[count] = color7
		case 8:
			randomcolorslice[count] = color8
		case 9:
			randomcolorslice[count] = color9
		case 0:
			randomcolorslice[count] = color10
		}

		count++

		if count == 10 {
			break
		}

	}

}
func drawaliencode() { // MARK: drawaliencode

	if aliencodeselected == false {
		selected := false
		for {
			choose := rInt(0, 10)
			checkaliencode := aliencodemap[choose]
			if checkaliencode.displayed == false {
				aliencodeactivenumber = choose
				checkaliencode.displayed = true
				aliencodemap[choose] = checkaliencode
				selected = true
			}
			if selected {
				break
			}
		}
		aliencodeselected = true
	}

	if aliencodeselected {
		if aliencodeactive {
			checkaliencode := aliencodemap[aliencodeactivenumber]
			rl.DrawTextureRec(imgs, checkaliencode.img, checkaliencode.v2, rl.Fade(rl.White, aliencodefade))
			checkaliencode.v2.Y += 5
			aliencodemap[aliencodeactivenumber] = checkaliencode
		}
	}

}
func drawpowerups() { // MARK:  drawpowerups

	if powerupactive {
		for a := 0; a < powerupnumber; a++ {
			checkpowerup := powerupsmap[a]
			if checkpowerup.active {
				rl.DrawCircle(int(checkpowerup.v2.X+28), int(checkpowerup.v2.Y+30), powerupcircleradius, rl.Fade(rl.Yellow, 0.3))
				rl.DrawTextureRec(imgs, checkpowerup.img, checkpowerup.v2, rl.White)

				if framecount%2 == 0 {
					powerupcircleradius += 4
					if powerupcircleradius >= 54 {
						powerupcircleradius = 20
					}
				}
			}
		}
	}
}
func drawfx() { // MARK: drawfx

	// snow
	if snowon && pauseon == false {

		for a := 0; a < len(snowtypemap); a++ {
			checksnow := snowv2map[a]
			snowtype := snowtypemap[a]

			if snowtype == 1 {
				rl.DrawTextureRec(imgs, snow1, checksnow, rl.Fade(rl.White, 0.7))
			} else if snowtype == 2 {
				rl.DrawTextureRec(imgs, snow2, checksnow, rl.Fade(rl.White, 0.5))
			} else if snowtype == 3 {
				rl.DrawTextureRec(imgs, snow3, checksnow, rl.Fade(rl.White, 0.6))
			} else if snowtype == 4 {
				rl.DrawTextureRec(imgs, snow4, checksnow, rl.Fade(rl.White, 0.8))
			}
			checksnow.Y += 7
			if flipcoin() {

				if flipcoin() {
					checksnow.X += 2
				} else {
					checksnow.X -= 2
				}

			}
			if checksnow.Y > float32(monh) {
				checksnow.Y = rFloat32(-40, -10)
				if checksnow.X > float32(monw) || checksnow.X < 0 {
					checksnow.X = rFloat32(10, monw-10)
				}
			}

			snowv2map[a] = checksnow

		}

	}
	// rain
	if rainon && pauseon == false {
		for a := 0; a < len(rainmap); a++ {
			checkrain := rainmap[a]
			rl.DrawCircle(int(checkrain.X), int(checkrain.Y), 2, rl.Fade(rl.Aqua, 0.8))
			checkrain.Y += 10
			if checkrain.Y > float32(monh) {
				checkrain.Y = rFloat32(-40, -10)
			}
			rainmap[a] = checkrain
		}
	}

	// night
	if nighton && pauseon == false {

		shadowv2 := rl.NewVector2(playerx-344, playery-362)
		rl.DrawTextureRec(imgs, nightshadow, shadowv2, rl.White)
		rl.DrawRectangle(0, 0, int(playerx-344), monh, rl.Fade(rl.Black, 0.8))                                       // left rec
		rl.DrawRectangle(int(playerx+456), 0, monw-int(playerx+456), monh, rl.Fade(rl.Black, 0.8))                   // right rec
		rl.DrawRectangle(int(playerx-344), 0, 800, int(playery-362), rl.Fade(rl.Black, 0.8))                         // top rec
		rl.DrawRectangle(int(playerx-344), int(playery+438), 800, (monh - int(playery+438)), rl.Fade(rl.Black, 0.8)) // bottom rec
	}

	// clouds
	if cloudson {
		drawclouds()
	}
	// scan lines
	if scanlineson {
		for a := 0; a < monh; a++ {
			rl.DrawLine(0, a, monw, a, rl.Fade(rl.Black, 0.1))
			a += 2
		}
	}

	// pixel noise
	if pixelnoiseon {
		if framecount%3 == 0 {
			for a := 0; a < 20; a++ {
				rl.DrawPixel(rInt(50, monw-50), rInt(50, monh-50), rl.Black)
				rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 2, 2, rl.Black)
			}
		}
	}
}
func drawcoins() { // MARK: drawcoins
	if pauseon == false {
		for a := 0; a < len(coinsmap); a++ {
			checkcoin := coinsmap[a]
			if checkcoin.active == false {
				if checkcoin.time == seconds {
					checkcoin.active = true
					coinsmap[a] = checkcoin
				}
			} else {
				if checkcoin.collected == false {
					rl.BeginMode2D(camera4X)
					rl.DrawTextureRec(imgs, coin, checkcoin.xy, rl.White)
					rl.EndMode2D()
					if checkcoin.xy.Y < float32((monh+100)/4) {
						checkcoin.xy.Y += 2
					}
					checkcoin.rec = rl.NewRectangle(checkcoin.xy.X*4, checkcoin.xy.Y*4, 64, 64)

					playercollisionrec5 := rl.NewRectangle(playerx, playery, 112, 75)
					if rl.CheckCollisionRecs(playercollisionrec5, checkcoin.rec) {
						coinnumber++
						checkcoin.active = false
						checkcoin.collected = true
					}

					coinsmap[a] = checkcoin
				}
			}
		}
	}
}
func drawplayer() { // MARK: drawplayer

	if playeroutlinecircleradiuson {
		if playeroutlinecircleradius > 20 {
			playeroutlinecircleradius -= 4
		} else {
			playeroutlinecircleradiuson = false
		}
	} else {
		if playeroutlinecircleradius < 60 {
			playeroutlinecircleradius += 4
		} else {
			playeroutlinecircleradiuson = true
		}

	}

	// draw player
	if playerdirection == 2 {
		if playership != ufo {
			rl.DrawCircle(int(playerxy.X+56), int(playerxy.Y+44), playeroutlinecircleradius, rl.Fade(scannercolor, 0.6))
		} else {
			rl.DrawCircle(int(playerxy.X+46), int(playerxy.Y+48), playeroutlinecircleradius+8, rl.Fade(scannercolor, 0.6))
		}
		playershadowv2 := rl.NewVector2(playerxy.X+4, playerxy.Y+8)
		rl.DrawTextureRec(imgs, playership, playershadowv2, rl.Black)
		rl.DrawTextureRec(imgs, playership, playerxy, rl.White)

		if playership != ufo {
			shipflamev2 := rl.NewVector2(playerxy.X+48, playerxy.Y+75)
			rl.DrawTextureRec(imgs, shipflame, shipflamev2, rl.Fade(rl.White, shipflamefade))
		}
	} else if playerdirection == 1 {
		if playership != ufo {
			rl.DrawCircle(int(playerxy.X), int(playerxy.Y), playeroutlinecircleradius, rl.Fade(scannercolor, 0.6))
		} else {
			rl.DrawCircle(int(playerxy.X), int(playerxy.Y), playeroutlinecircleradius+8, rl.Fade(scannercolor, 0.6))
		}
		shipleft := rl.NewRectangle(playerxy.X, playerxy.Y, 112, 75)
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, shipleft, originv2, -15, rl.White)
	} else if playerdirection == 3 {
		if playership != ufo {
			rl.DrawCircle(int(playerxy.X), int(playerxy.Y), playeroutlinecircleradius, rl.Fade(scannercolor, 0.6))
		} else {
			rl.DrawCircle(int(playerxy.X), int(playerxy.Y), playeroutlinecircleradius+8, rl.Fade(scannercolor, 0.6))
		}
		shipleft := rl.NewRectangle(playerxy.X, playerxy.Y, 112, 75)
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, shipleft, originv2, 15, rl.White)
	}

	// draw laser marker
	if currenttarget == lasermarker5 {
		lasermarkerv2 := rl.NewVector2(playerxy.X+24, 200)
		lasermarkershadowv2 := rl.NewVector2(playerxy.X+26, 202)
		if playerxy.Y < 250 {
			lasermarkerv2.Y = 0
			lasermarkershadowv2.Y = 0
		}
		rl.DrawTextureRec(imgs, currenttarget, lasermarkershadowv2, rl.Black)
		rl.DrawTextureRec(imgs, currenttarget, lasermarkerv2, rl.White)
	} else {
		lasermarkerv2 := rl.NewVector2(playerxy.X+36, 200)
		lasermarkershadowv2 := rl.NewVector2(playerxy.X+38, 202)
		if playerxy.Y < 250 {
			lasermarkerv2.Y = 0
			lasermarkershadowv2.Y = 0
		}
		rl.DrawTextureRec(imgs, currenttarget, lasermarkershadowv2, rl.Black)
		rl.DrawTextureRec(imgs, currenttarget, lasermarkerv2, rl.White)
	}

}
func drawclouds() { // MARK: drawclouds
	rl.BeginMode2D(cameraclouds)
	cloud1shadowv2 := rl.NewVector2(cloud1v2.X, cloud1v2.Y+150)
	cloud2shadowv2 := rl.NewVector2(cloud2v2.X, cloud2v2.Y+150)
	cloud3shadowv2 := rl.NewVector2(cloud3v2.X, cloud3v2.Y+150)

	rl.DrawTextureRec(imgs, cloud1, cloud1shadowv2, rl.Fade(rl.Black, 0.1))
	rl.DrawTextureRec(imgs, cloud1, cloud1v2, cloudcolor)
	rl.DrawTextureRec(imgs, cloud2, cloud2shadowv2, rl.Fade(rl.Black, 0.1))
	rl.DrawTextureRec(imgs, cloud2, cloud2v2, cloudcolor)
	rl.DrawTextureRec(imgs, cloud3, cloud3shadowv2, rl.Fade(rl.Black, 0.1))
	rl.DrawTextureRec(imgs, cloud3, cloud3v2, cloudcolor)
	rl.EndMode2D()
}
func drawminiboss() { // MARK: drawminiboss

	if createminibosscomplete == false {
		minibossmodifiers.minibosstype = rInt(1, 4)

		minibossmodifiers.hp = rInt(10, 20)
		minibossmodifiers.speed = rFloat32(5, 11)
		if minibossmodifiers.minibosstype == 1 {
			minibossmodifiers.firetype = rInt(1, 4)
			minibossmodifiers.v2.X = rFloat32(monw/4, monw-(monw/4))
			minibossmodifiers.v2.Y = rFloat32(-600, -450)
		} else if minibossmodifiers.minibosstype == 2 {
			minibossmodifiers.firetype = 4
			minibossmodifiers.v2.X = rFloat32(-500, -350)
			minibossmodifiers.v2.Y = rFloat32(10, monh-300)
		} else if minibossmodifiers.minibosstype == 3 {
			minibossmodifiers.firetype = 5
			minibossmodifiers.v2.X = rFloat32(50, 300)
			minibossmodifiers.v2.Y = rFloat32(-300, -200)
			minibossmodifiers.v22.X = rFloat32(400, 700)
			minibossmodifiers.v22.Y = rFloat32(-300, -200)
			minibossmodifiers.v23.X = rFloat32(800, monw-300)
			minibossmodifiers.v23.Y = rFloat32(-300, -200)
		}
		createminibosscomplete = true
	}

	// draw miniboss bullets
	for a := 0; a < len(minibossbulletmap); a++ {
		checkbullet := minibossbulletmap[a]
		if minibossmodifiers.minibosstype == 1 {
			rl.DrawCircle(int(checkbullet.v2.X-2), int(checkbullet.v2.Y-2), 12, rl.Black)
			rl.DrawCircle(int(checkbullet.v2.X), int(checkbullet.v2.Y), 12, minibossbulletcolor)
			rl.DrawCircle(int(checkbullet.v2.X), int(checkbullet.v2.Y), 4, rl.Black)
			switch checkbullet.direction {
			case 1:
				checkbullet.v2.X -= 8
				checkbullet.v2.Y -= 8
			case 2:
				checkbullet.v2.X += 8
				checkbullet.v2.Y -= 8
			case 3:
				checkbullet.v2.X += 8
				checkbullet.v2.Y += 12
			case 4:
				checkbullet.v2.X -= 8
				checkbullet.v2.Y += 12
			case 5:
				checkbullet.v2.Y -= 12
			case 6:
				checkbullet.v2.X += 12
				checkbullet.v2.Y += minibossmodifiers.speed
			case 7:
				checkbullet.v2.Y += 12
			case 8:
				checkbullet.v2.X -= 12
				checkbullet.v2.Y += minibossmodifiers.speed
			}
			minibossbulletmap[a] = checkbullet
		} else if minibossmodifiers.minibosstype == 2 {
			rl.DrawTextureRec(imgs, checkbullet.img, checkbullet.v2, rl.White)
			switch checkbullet.direction {
			case 5:
				checkbullet.v2.Y -= 12
				if flipcoin() {
					checkbullet.v2.X += rFloat32(-8, 9)
				}
			case 7:
				checkbullet.v2.Y += 12
				if flipcoin() {
					checkbullet.v2.X += rFloat32(-8, 9)
				}
			}
			minibossbulletmap[a] = checkbullet
		} else if minibossmodifiers.minibosstype == 3 {
			rl.DrawTextureRec(imgs, checkbullet.img, checkbullet.v2, rl.White)
			switch checkbullet.direction {
			case 7:
				checkbullet.v2.Y += 18
			}
			minibossbulletmap[a] = checkbullet
		}
	}

	if framecount%6 == 0 {
		if minibossbulletcolor == rl.Purple {
			minibossbulletcolor = rl.Red
		} else if minibossbulletcolor == rl.Red {
			minibossbulletcolor = rl.Pink
		} else if minibossbulletcolor == rl.Pink {
			minibossbulletcolor = rl.Orange
		} else if minibossbulletcolor == rl.Orange {
			minibossbulletcolor = rl.Yellow
		} else if minibossbulletcolor == rl.Yellow {
			minibossbulletcolor = rl.Purple
		}

	}

	bulletychange := float32(0)

	switch minibossmodifiers.firetype {

	case 5:
		if framecount%15 == 0 {
			newbullet := minibossbullet{}
			newbullet.active = true
			newbullet.direction = 7
			choose := rInt(1, 7)
			switch choose {
			case 1:
				newbullet.img = missile7
				newbullet.hp = 2
			case 2:
				newbullet.img = missile8
				newbullet.hp = 3
			case 3:
				newbullet.img = missile9
				newbullet.hp = 4
			case 4:
				newbullet.img = missile10
				newbullet.hp = 5
			case 5:
				newbullet.img = missile11
				newbullet.hp = 6
			case 6:
				newbullet.img = missile12
				newbullet.hp = 7
			}
			newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+80, minibossmodifiers.v2.Y+140)
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++

			choose = rInt(1, 7)
			switch choose {
			case 1:
				newbullet.img = missile7
				newbullet.hp = 2
			case 2:
				newbullet.img = missile8
				newbullet.hp = 3
			case 3:
				newbullet.img = missile9
				newbullet.hp = 4
			case 4:
				newbullet.img = missile10
				newbullet.hp = 5
			case 5:
				newbullet.img = missile11
				newbullet.hp = 6
			case 6:
				newbullet.img = missile12
				newbullet.hp = 7
			}
			newbullet.v2 = rl.NewVector2(minibossmodifiers.v22.X+80, minibossmodifiers.v22.Y+140)
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++

			choose = rInt(1, 7)
			switch choose {
			case 1:
				newbullet.img = missile7
				newbullet.hp = 2
			case 2:
				newbullet.img = missile8
				newbullet.hp = 3
			case 3:
				newbullet.img = missile9
				newbullet.hp = 4
			case 4:
				newbullet.img = missile10
				newbullet.hp = 5
			case 5:
				newbullet.img = missile11
				newbullet.hp = 6
			case 6:
				newbullet.img = missile12
				newbullet.hp = 7
			}
			newbullet.v2 = rl.NewVector2(minibossmodifiers.v23.X+80, minibossmodifiers.v23.Y+140)
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
		}
	case 4:
		if framecount%6 == 0 {
			newbullet := minibossbullet{}

			newbullet.active = true

			if flipcoin() {
				newbullet.direction = 5
				choose := rInt(1, 7)
				switch choose {
				case 1:
					newbullet.img = missile1
					newbullet.hp = 2
				case 2:
					newbullet.img = missile2
					newbullet.hp = 3
				case 3:
					newbullet.img = missile3
					newbullet.hp = 4
				case 4:
					newbullet.img = missile4
					newbullet.hp = 5
				case 5:
					newbullet.img = missile5
					newbullet.hp = 6
				case 6:
					newbullet.img = missile6
					newbullet.hp = 7
				}

			} else {
				newbullet.direction = 7
				choose := rInt(1, 7)
				switch choose {
				case 1:
					newbullet.img = missile7
					newbullet.hp = 2
				case 2:
					newbullet.img = missile8
					newbullet.hp = 3
				case 3:
					newbullet.img = missile9
					newbullet.hp = 4
				case 4:
					newbullet.img = missile10
					newbullet.hp = 5
				case 5:
					newbullet.img = missile11
					newbullet.hp = 6
				case 6:
					newbullet.img = missile12
					newbullet.hp = 7
				}

			}
			newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+140, minibossmodifiers.v2.Y+90)
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++

		}
		if minibossbulletcount > 480 {
			minibossbulletcount = 0
		}

	case 3:
		if framecount%15 == 0 {
			for a := 0; a < 10; a++ {
				newbullet := minibossbullet{}
				newbullet.direction = 6
				newbullet.hp = 4
				newbullet.active = true
				if minibossmodifiers.minibosstype == 1 {
					newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, (minibossmodifiers.v2.Y+400)-bulletychange)
				} else if minibossmodifiers.minibosstype == 2 {
					newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
				} else if minibossmodifiers.minibosstype == 3 {
					newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
				}
				minibossbulletmap[minibossbulletcount] = newbullet
				minibossbulletcount++
				newbullet.direction = 8
				newbullet.hp = 4
				newbullet.active = true
				if minibossmodifiers.minibosstype == 1 {
					newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, (minibossmodifiers.v2.Y+400)-bulletychange)
				} else if minibossmodifiers.minibosstype == 2 {
					newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
				} else if minibossmodifiers.minibosstype == 3 {
					newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
				}
				minibossbulletmap[minibossbulletcount] = newbullet
				minibossbulletcount++
				bulletychange += 44
			}
		}
		if minibossbulletcount > 480 {
			minibossbulletcount = 0
		}
	case 2:
		if framecount%15 == 0 {
			newbullet := minibossbullet{}
			newbullet.direction = 5
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
			newbullet.direction = 6
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
			newbullet.direction = 7
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
			newbullet.direction = 8
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
		}

		if minibossbulletcount > 480 {
			minibossbulletcount = 0
		}
	case 1:
		if framecount%15 == 0 {
			newbullet := minibossbullet{}
			newbullet.direction = 1
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
			newbullet.direction = 2
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
			newbullet.direction = 3
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
			newbullet.direction = 4
			newbullet.hp = 4
			newbullet.active = true
			if minibossmodifiers.minibosstype == 1 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 2 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			} else if minibossmodifiers.minibosstype == 3 {
				newbullet.v2 = rl.NewVector2(minibossmodifiers.v2.X+142, minibossmodifiers.v2.Y+320)
			}
			minibossbulletmap[minibossbulletcount] = newbullet
			minibossbulletcount++
		}

		if minibossbulletcount > 480 {
			minibossbulletcount = 0
		}

	}

	// draw miniboss image
	switch minibossmodifiers.minibosstype {
	case 1:
		shadowv2 := rl.NewVector2(minibossmodifiers.v2.X+4, minibossmodifiers.v2.Y+4)
		rl.DrawTextureRec(imgs, miniboss1, shadowv2, rl.Black)
		rl.DrawTextureRec(imgs, miniboss1, minibossmodifiers.v2, rl.White)
	case 2:
		shadowv2 := rl.NewVector2(minibossmodifiers.v2.X+4, minibossmodifiers.v2.Y+4)
		rl.DrawTextureRec(imgs, miniboss2, shadowv2, rl.Black)
		rl.DrawTextureRec(imgs, miniboss2, minibossmodifiers.v2, rl.White)
	case 3:
		shadowv2 := rl.NewVector2(minibossmodifiers.v2.X+4, minibossmodifiers.v2.Y+4)
		rl.DrawTextureRec(imgs, miniboss3, shadowv2, rl.Black)
		rl.DrawTextureRec(imgs, miniboss3, minibossmodifiers.v2, rl.White)
		shadowv2 = rl.NewVector2(minibossmodifiers.v22.X+4, minibossmodifiers.v22.Y+4)
		rl.DrawTextureRec(imgs, miniboss3, shadowv2, rl.Black)
		rl.DrawTextureRec(imgs, miniboss3, minibossmodifiers.v22, rl.White)
		shadowv2 = rl.NewVector2(minibossmodifiers.v23.X+4, minibossmodifiers.v23.Y+4)
		rl.DrawTextureRec(imgs, miniboss3, shadowv2, rl.Black)
		rl.DrawTextureRec(imgs, miniboss3, minibossmodifiers.v23, rl.White)
	}
	// move miniboss image
	if minibossmodifiers.minibosstype == 1 {
		if minibossmodifiers.v2.Y < float32(monh+10) {
			minibossmodifiers.v2.Y += minibossmodifiers.speed
		} else {
			createminibosscomplete = false
			bulletychange = 0
			minibossbulletcount = 0
			for a := 0; a < len(minibossbulletmap); a++ {
				minibossbulletmap[a] = minibossbullet{}
			}
		}
	} else if minibossmodifiers.minibosstype == 2 {
		if minibossmodifiers.v2.X < float32(monw+10) {
			minibossmodifiers.v2.X += minibossmodifiers.speed
		} else {
			createminibosscomplete = false
			bulletychange = 0
			minibossbulletcount = 0
			for a := 0; a < len(minibossbulletmap); a++ {
				minibossbulletmap[a] = minibossbullet{}
			}
		}
	} else if minibossmodifiers.minibosstype == 3 {
		if minibossmodifiers.v2.Y < float32(monh+10) || minibossmodifiers.v22.Y < float32(monh+10) || minibossmodifiers.v23.Y < float32(monh+10) {
			minibossmodifiers.v2.Y += minibossmodifiers.speed
			minibossmodifiers.v22.Y += minibossmodifiers.speed
			minibossmodifiers.v23.Y += minibossmodifiers.speed
		} else {
			createminibosscomplete = false
			bulletychange = 0
			minibossbulletcount = 0
			for a := 0; a < len(minibossbulletmap); a++ {
				minibossbulletmap[a] = minibossbullet{}
			}
		}
	}

}
func drawenemies() { // MARK: drawenemies

	// draw enemies
	for a := 0; a < len(enemymap); a++ {
		checkenemy := enemymap[a]
		if checkenemy.hp > 0 {
			enemyshadowv2 := rl.NewVector2(checkenemy.xy.X+5, checkenemy.xy.Y-5)
			rl.DrawTextureRec(imgs, checkenemy.img, enemyshadowv2, rl.Black)
			rl.DrawTextureRec(imgs, checkenemy.img, checkenemy.xy, rl.White)
		}
	}

}
func drawbullets() { // MARK: drawbullets
	// draw bullets
	for a := 0; a < len(bulletmap); a++ {
		checkbullet := bulletmap[a]
		if checkbullet.active {
			shadowv2 := rl.NewVector2(checkbullet.xy.X+2, checkbullet.xy.Y+2)
			switch bulletmodifiers.img {
			case 1:
				rl.DrawTextureRec(imgs, bullet1, shadowv2, rl.Black)
				rl.DrawTextureRec(imgs, bullet1, checkbullet.xy, rl.White)
			case 2:
				rl.DrawTextureRec(imgs, bullet2, shadowv2, rl.Black)
				rl.DrawTextureRec(imgs, bullet2, checkbullet.xy, rl.White)
			case 3:
				rl.DrawTextureRec(imgs, bullet3, shadowv2, rl.Black)
				rl.DrawTextureRec(imgs, bullet3, checkbullet.xy, rl.White)
			case 4:
				rl.DrawTextureRec(imgs, bullet4, shadowv2, rl.Black)
				rl.DrawTextureRec(imgs, bullet4, checkbullet.xy, rl.White)

			}
		}
	}
}
func drawinfobar() { // MARK: drawinfobar

	// playerhp bar
	rl.DrawText("hp", 67, monh-137, 40, rl.Black)
	rl.DrawText("hp", 69, monh-139, 40, rl.White)
	rl.DrawText("hp", 70, monh-140, 40, rl.Black)
	rl.DrawRectangle(50, monh-90, 302, 70, rl.Maroon)
	rl.DrawRectangleLines(50, monh-90, 302, 70, rl.Black)
	rl.DrawRectangleLines(49, monh-91, 304, 72, rl.Black)

	xchange := 0
	for a := 0; a < playerhp; a++ {
		rl.DrawRectangle(51+xchange, monh-89, 3, 68, rl.Color(green1))
		xchange += 3
	}
	playerhpTEXT := strconv.Itoa(playerhp)
	rl.DrawText(playerhpTEXT, 177, monh-72, 40, rl.Black)
	rl.DrawText(playerhpTEXT, 179, monh-74, 40, rl.White)
	rl.DrawText(playerhpTEXT, 180, monh-75, 40, rl.Black)

	// draw powerups
	rl.DrawText("powerups", 427, monh-137, 40, rl.Black)
	rl.DrawText("powerups", 429, monh-139, 40, rl.White)
	rl.DrawText("powerups", 430, monh-140, 40, rl.Black)
	xchangepowerups := 0
	for a := 0; a < 5; a++ {
		rl.DrawRectangleLines(400+xchangepowerups, monh-91, 72, 72, rl.Black)
		rl.DrawRectangleLines(401+xchangepowerups, monh-90, 70, 70, rl.Black)
		rl.DrawRectangle(401+xchangepowerups, monh-90, 70, 70, rl.Fade(rl.Black, 0.5))
		xchangepowerups += 76
	}
	powerimgsv2 := rl.NewVector2(410, float32(monh-84))
	for a := 0; a < len(playerpowerups); a++ {
		checkpowerup := playerpowerups[a]
		if checkpowerup.active {
			rl.DrawTextureRec(imgs, checkpowerup.img, powerimgsv2, rl.White)
			powerimgsv2.X += 76
		}
	}
	// draw coin
	coinv2 := rl.NewVector2(202, float32((monh-88)/4))
	rl.BeginMode2D(camera4X)
	rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	rl.EndMode2D()
	coinnumberTEXT := strconv.Itoa(coinnumber)
	rl.DrawText(coinnumberTEXT, 881, monh-91, 80, rl.Black)
	rl.DrawText(coinnumberTEXT, 882, monh-92, 80, scorecolor)
	rl.DrawText(coinnumberTEXT, 884, monh-94, 80, rl.White)
	rl.DrawText(coinnumberTEXT, 885, monh-95, 80, rl.Black)
	// draw alien code
	aliencodev2 := rl.NewVector2(1000, float32(monh-110))
	for a := 0; a < len(aliencodemapplayer); a++ {
		checkaliencode := aliencodemapplayer[a]
		rl.DrawTextureRec(imgs, checkaliencode.img, aliencodev2, rl.White)
		aliencodev2.X += 80
	}

	// draw score
	scoreTEXT := strconv.Itoa(totalscore)
	rl.DrawText("score", monw-223, monh-171, 40, rl.Black)
	rl.DrawText("score", monw-221, monh-173, 40, rl.White)
	rl.DrawText("score", monw-220, monh-174, 40, rl.Black)

	rl.DrawText(scoreTEXT, monw-(185+xchangescoretext), monh-131, 140, rl.Black)
	rl.DrawText(scoreTEXT, monw-(183+xchangescoretext), monh-133, 140, scorecolor)
	rl.DrawText(scoreTEXT, monw-(181+xchangescoretext), monh-135, 140, rl.White)
	rl.DrawText(scoreTEXT, monw-(180+xchangescoretext), monh-136, 140, rl.Black)

	if totalscore > 9 && totalscore < 100 {
		xchangescoretext = 50
	}
	if totalscore >= 100 && totalscore < 1000 {
		xchangescoretext = 100
	}
	if totalscore >= 1000 && totalscore < 10000 {
		xchangescoretext = 150
	}
	if totalscore >= 10000 && totalscore < 100000 {
		xchangescoretext = 200
	}

	// timer
	secondsTEXT := strconv.Itoa(seconds)
	rl.DrawText(secondsTEXT, 49, 51, 40, rl.White)
	rl.DrawText(secondsTEXT, 50, 50, 40, rl.Black)

	/*
		// timer
			secondsTEXT := strconv.Itoa(seconds)
			rl.DrawText(secondsTEXT, 49, 51, 40, rl.White)
			rl.DrawText(secondsTEXT, 50, 50, 40, rl.Black)

		// level name
			rl.DrawText(currentlevelname, 119, 51, 40, rl.White)
			rl.DrawText(currentlevelname, 120, 50, 40, rl.Black)

			// draw mods
			rl.DrawText("mods", 847, monh-137, 40, rl.Black)
			rl.DrawText("mods", 849, monh-139, 40, rl.White)
			rl.DrawText("mods", 850, monh-140, 40, rl.Black)
			xchangemods := 0
			for a := 0; a < 3; a++ {
				rl.DrawRectangleLines(820+xchangemods, monh-91, 72, 72, rl.Black)
				rl.DrawRectangleLines(821+xchangemods, monh-90, 70, 70, rl.Black)
				rl.DrawRectangle(821+xchangemods, monh-90, 70, 70, rl.Fade(rl.Black, 0.5))
				xchangemods += 76
			}

			// draw items
			rl.DrawText("items", 1117, monh-137, 40, rl.Black)
			rl.DrawText("items", 1119, monh-139, 40, rl.White)
			rl.DrawText("items", 1120, monh-140, 40, rl.Black)
			xchangeitems := 0
			for a := 0; a < 3; a++ {
				rl.DrawRectangleLines(1088+xchangeitems, monh-91, 72, 72, rl.Black)
				rl.DrawRectangleLines(1089+xchangeitems, monh-90, 70, 70, rl.Black)
				rl.DrawRectangle(1089+xchangeitems, monh-90, 70, 70, rl.Fade(rl.Black, 0.5))
				xchangeitems += 76
			}

	*/

}
func drawbosslevel() { // MARK: drawbosslevel
	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Black, bossbackfade))
	rl.BeginMode2D(camera)

	// draw visible blocks
	drawblock = nextblock
	for a := 0; a < screenblocknumber; a++ {
		checkblock := blockmap[drawblock]

		rl.DrawLineV(checkblock.xy, checkblock.xy2, rl.Fade(rl.DarkBlue, bosslinesfade))
		rl.DrawLineV(checkblock.xy2, checkblock.xy3, rl.Fade(rl.DarkBlue, bosslinesfade))
		rl.DrawLineV(checkblock.xy, checkblock.xy4, rl.Fade(rl.DarkBlue, bosslinesfade))
		rl.DrawLineV(checkblock.xy4, checkblock.xy3, rl.Fade(rl.DarkBlue, bosslinesfade))

		if drawblock == highlightedtileblock1 {
			rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Red, highlightfade1))
			rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Red, highlightfade1))
		}
		if drawblock == highlightedtileblock2 {
			rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Green, highlightfade2))
			rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Green, highlightfade2))
		}
		if drawblock == highlightedtileblock3 {
			rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.Blue, highlightfade3))
			rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.Blue, highlightfade3))
		}
		if drawblock == highlightedtileblock4 {
			rl.DrawTriangle(checkblock.xy2, checkblock.xy, checkblock.xy4, rl.Fade(rl.White, highlightfade4))
			rl.DrawTriangle(checkblock.xy2, checkblock.xy4, checkblock.xy3, rl.Fade(rl.White, highlightfade4))
		}
		drawblock++
	}

	rl.EndMode2D()
	rl.DrawRectangleGradientV(0, monh-80, monw, 80, rl.Black, rl.DarkBlue)
	rl.DrawRectangleGradientV(0, monh-160, monw, 80, rl.Transparent, rl.Black)

	bossv2 := rl.NewVector2(float32(monw/3), 10)
	rl.DrawTextureRec(imgs, boss1, bossv2, rl.White)

}
func drawshop() { // MARK: drawshop
	cloudson = false
	pauseon = true
	stoptimer = true

	if framecount%60 == 0 {
		if spacebaroff {
			spacebaroff = false
		}
	}

	rl.DrawRectangle(0, 0, monw, monh, rl.Black)

	if framecount%3 == 0 {
		for a := 0; a < 20; a++ {
			randomcolors()
			pixelcolor1 := randomcolorslice[0]
			rl.DrawPixel(rInt(50, monw-50), rInt(50, monh-50), pixelcolor1)
			reccolor := randomcolorslice[1]
			rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 3, 3, reccolor)
			reccolor2 := randomcolorslice[2]
			rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 2, 2, reccolor2)
		}
	}

	rl.DrawText("the shop", 80-3, 50+3, 80, rl.Maroon)
	rl.DrawText("the shop", 80-1, 50+1, 80, rl.Black)
	rl.DrawText("the shop", 80, 50, 80, rl.White)
	// drawcoin number
	rl.BeginMode2D(camera4X)
	coin4xv2 := rl.NewVector2(120, 15)
	rl.DrawTextureRec(imgs, coin, coin4xv2, rl.White)
	rl.EndMode2D()
	coinsTEXT := strconv.Itoa(coinnumber)
	rl.DrawText("x", 560-3, 50+3, 80, rl.Maroon)
	rl.DrawText("x", 560-1, 50+1, 80, rl.Black)
	rl.DrawText("x", 560, 50, 80, rl.White)
	rl.DrawText(coinsTEXT, 620-3, 50+3, 80, rl.Maroon)
	rl.DrawText(coinsTEXT, 620-1, 50+1, 80, rl.Black)
	rl.DrawText(coinsTEXT, 620, 50, 80, rl.White)
	// exit button

	if storeselect == 0 {
		rl.DrawRectangle(994, 54, 232, 72, rl.Fade(rl.DarkGray, 0.3))
		rl.DrawRectangleLines(994, 54, 232, 72, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(993, 53, 234, 74, rl.Fade(rl.Green, storeselectfade))
	}

	rl.DrawRectangle(1000, 60, 220, 60, rl.Maroon)
	rl.DrawText("exit", 1070-2, 66+2, 50, rl.Black)
	rl.DrawText("exit", 1070, 66, 50, rl.White)

	// select box
	storeselectv2 := rl.NewVector2(115, 195)
	storex := int(storeselectv2.X)
	storey := int(storeselectv2.Y)
	if storeselect == 1 { // ship box 1
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 2 {
		storex += 170
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Yellow, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Yellow, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 3 {
		storex += 340
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Red, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Red, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 4 {
		storex += 500
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Blue, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Blue, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 5 { // laser target box 1
		storex += 650
		rl.DrawRectangleLines(storex-1, storey-1, 132, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 130, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 130, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 6 {
		storex += 785
		rl.DrawRectangleLines(storex-1, storey-1, 132, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 130, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 130, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 7 {
		storex += 924
		rl.DrawRectangleLines(storex-1, storey-1, 132, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 130, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 130, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 8 {
		storex += 1062
		rl.DrawRectangleLines(storex-1, storey-1, 132, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 130, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 130, 183, rl.Fade(rl.DarkGray, 0.2))
	}
	storex = int(storeselectv2.X)
	storey = 400
	if storeselect == 9 {
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 10 {
		storex += 170
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Yellow, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Yellow, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 11 {
		storex += 340
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Red, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Red, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 12 {
		storex += 500
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Purple, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Purple, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 13 {
		storex += 650
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 14 {
		storex += 785
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 15 {
		storex += 924
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	} else if storeselect == 16 {
		storex += 1062
		rl.DrawRectangleLines(storex-1, storey-1, 144, 185, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangleLines(storex, storey, 142, 183, rl.Fade(rl.Green, storeselectfade))
		rl.DrawRectangle(storex, storey, 142, 183, rl.Fade(rl.DarkGray, 0.2))
	}
	// shop scanners
	row2v2 := rl.NewVector2(185, 470)
	rl.DrawCircle(int(row2v2.X), int(row2v2.Y), shopscannerradius2, rl.Fade(rl.DarkGreen, 0.5))
	if !greenscannerpurchased {
		rl.DrawText("x5", int(row2v2.X-6), int(row2v2.Y+70), 30, rl.White)
	} else {
		if scannercolor == rl.DarkGreen {
			rl.DrawText("unequip", int(row2v2.X-54), int(row2v2.Y+72), 30, rl.Blue)
			rl.DrawText("unequip", int(row2v2.X-52), int(row2v2.Y+70), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row2v2.X-37), int(row2v2.Y+72), 30, rl.Magenta)
			rl.DrawText("equip", int(row2v2.X-35), int(row2v2.Y+70), 30, rl.White)
		}
	}
	row2v2.X += 170
	rl.DrawCircle(int(row2v2.X), int(row2v2.Y), shopscannerradius2, rl.Fade(rl.Yellow, 0.5))
	if !yellowscannerpurchased {
		rl.DrawText("x5", int(row2v2.X-6), int(row2v2.Y+70), 30, rl.White)
	} else {
		if scannercolor == rl.Yellow {
			rl.DrawText("unequip", int(row2v2.X-54), int(row2v2.Y+72), 30, rl.Blue)
			rl.DrawText("unequip", int(row2v2.X-52), int(row2v2.Y+70), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row2v2.X-37), int(row2v2.Y+72), 30, rl.Magenta)
			rl.DrawText("equip", int(row2v2.X-35), int(row2v2.Y+70), 30, rl.White)
		}
	}
	row2v2.X += 170
	rl.DrawCircle(int(row2v2.X), int(row2v2.Y), shopscannerradius2, rl.Fade(rl.Red, 0.5))
	if !redscannerpurchashed {
		rl.DrawText("x5", int(row2v2.X-6), int(row2v2.Y+70), 30, rl.White)
	} else {
		if scannercolor == rl.Red {
			rl.DrawText("unequip", int(row2v2.X-54), int(row2v2.Y+72), 30, rl.Blue)
			rl.DrawText("unequip", int(row2v2.X-52), int(row2v2.Y+70), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row2v2.X-37), int(row2v2.Y+72), 30, rl.Magenta)
			rl.DrawText("equip", int(row2v2.X-35), int(row2v2.Y+70), 30, rl.White)
		}
	}
	row2v2.X += 162
	rl.DrawCircle(int(row2v2.X), int(row2v2.Y), shopscannerradius2, rl.Fade(rl.Purple, 0.5))
	if !purplescannerpurchased {
		rl.DrawText("x5", int(row2v2.X-10), int(row2v2.Y+70), 30, rl.White)
	} else {
		if scannercolor == rl.Purple {
			rl.DrawText("unequip", int(row2v2.X-54), int(row2v2.Y+72), 30, rl.Blue)
			rl.DrawText("unequip", int(row2v2.X-52), int(row2v2.Y+70), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row2v2.X-37), int(row2v2.Y+72), 30, rl.Magenta)
			rl.DrawText("equip", int(row2v2.X-35), int(row2v2.Y+70), 30, rl.White)
		}
	}
	// shop ships
	row1v2 := rl.NewVector2(130, 210)
	rl.DrawTextureRec(imgs, ship2, row1v2, rl.White)
	rl.DrawText("-1 damage", int(row1v2.X+6), int(row1v2.Y+90), 20, rl.White)
	if !greenshippurchased {
		rl.DrawText("x5", int(row1v2.X+50), int(row1v2.Y+128), 30, rl.White)
	} else {
		if playership == ship2 {
			rl.DrawText("unequip", int(row1v2.X-2), int(row1v2.Y+130), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X), int(row1v2.Y+128), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X+16), int(row1v2.Y+130), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X+18), int(row1v2.Y+128), 30, rl.White)
		}
	}
	row1v2.X += 170
	rl.DrawTextureRec(imgs, ship3, row1v2, rl.White)
	rl.DrawText("+1 attack", int(row1v2.X+6), int(row1v2.Y+90), 20, rl.White)
	if !yellowshippurchased {
		rl.DrawText("x5", int(row1v2.X+50), int(row1v2.Y+128), 30, rl.White)
	} else {
		if playership == ship3 {
			rl.DrawText("unequip", int(row1v2.X-2), int(row1v2.Y+130), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X), int(row1v2.Y+128), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X+16), int(row1v2.Y+130), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X+18), int(row1v2.Y+128), 30, rl.White)
		}
	}
	row1v2.X += 170
	rl.DrawTextureRec(imgs, ship4, row1v2, rl.White)
	rl.DrawText("+20% speed", int(row1v2.X-4), int(row1v2.Y+90), 20, rl.White)
	if !redshippurchased {
		rl.DrawText("x5", int(row1v2.X+50), int(row1v2.Y+128), 30, rl.White)
	} else {
		if playership == ship4 {
			rl.DrawText("unequip", int(row1v2.X-2), int(row1v2.Y+130), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X), int(row1v2.Y+128), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X+16), int(row1v2.Y+130), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X+18), int(row1v2.Y+128), 30, rl.White)
		}
	}
	row1v2.X += 170
	row1v2.Y -= 2
	rl.DrawTextureRec(imgs, ufo, row1v2, rl.White)
	rl.DrawText("ufo++", int(row1v2.X+20), int(row1v2.Y+100), 20, rl.White)
	if !ufopurchased {
		rl.DrawText("x50", int(row1v2.X+36), int(row1v2.Y+130), 30, rl.White)
	} else {
		if playership == ufo {
			rl.DrawText("unequip", int(row1v2.X-10), int(row1v2.Y+130), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X-8), int(row1v2.Y+128), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X+6), int(row1v2.Y+130), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X+8), int(row1v2.Y+128), 30, rl.White)
		}
	}
	row1v2.X += 162
	row1v2.Y += 28

	// shop laser targets
	lastertargetv2shadow := rl.NewVector2(row1v2.X-2, row1v2.Y+2)
	rl.DrawTextureRec(imgs, lasermarker5, lastertargetv2shadow, rl.DarkGray)
	rl.DrawTextureRec(imgs, lasermarker5, row1v2, rl.White)
	if !target1purchased {
		rl.DrawText("x5", int(row1v2.X+18), int(row1v2.Y+102), 30, rl.White)
	} else {
		if currenttarget == lasermarker5 {
			rl.DrawText("unequip", int(row1v2.X-28), int(row1v2.Y+102), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X-26), int(row1v2.Y+100), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X-18), int(row1v2.Y+102), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X-16), int(row1v2.Y+100), 30, rl.White)
		}
	}
	row1v2.X += 144
	lastertargetv2shadow.X += 144
	rl.DrawTextureRec(imgs, lasermarker2, lastertargetv2shadow, rl.DarkGray)
	rl.DrawTextureRec(imgs, lasermarker2, row1v2, rl.White)
	if !target2purchased {
		rl.DrawText("x5", int(row1v2.X+16), int(row1v2.Y+102), 30, rl.White)
	} else {
		if currenttarget == lasermarker2 {
			rl.DrawText("unequip", int(row1v2.X-38), int(row1v2.Y+102), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X-36), int(row1v2.Y+100), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X-18), int(row1v2.Y+102), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X-16), int(row1v2.Y+100), 30, rl.White)
		}
	}
	row1v2.X += 136
	lastertargetv2shadow.X += 140
	rl.DrawTextureRec(imgs, lasermarker3, lastertargetv2shadow, rl.DarkGray)
	rl.DrawTextureRec(imgs, lasermarker3, row1v2, rl.White)
	if !target3purchased {
		rl.DrawText("x5", int(row1v2.X+18), int(row1v2.Y+102), 30, rl.White)
	} else {
		if currenttarget == lasermarker3 {
			rl.DrawText("unequip", int(row1v2.X-34), int(row1v2.Y+102), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X-32), int(row1v2.Y+100), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X-18), int(row1v2.Y+102), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X-16), int(row1v2.Y+100), 30, rl.White)
		}
	}
	row1v2.X += 138
	lastertargetv2shadow.X += 134
	rl.DrawTextureRec(imgs, lasermarker4, lastertargetv2shadow, rl.DarkGray)
	rl.DrawTextureRec(imgs, lasermarker4, row1v2, rl.White)
	if !target4purchased {
		rl.DrawText("x5", int(row1v2.X+22), int(row1v2.Y+102), 30, rl.White)
	} else {
		if currenttarget == lasermarker4 {
			rl.DrawText("unequip", int(row1v2.X-34), int(row1v2.Y+102), 30, rl.Blue)
			rl.DrawText("unequip", int(row1v2.X-32), int(row1v2.Y+100), 30, rl.White)
		} else {
			rl.DrawText("equip", int(row1v2.X-18), int(row1v2.Y+102), 30, rl.Magenta)
			rl.DrawText("equip", int(row1v2.X-16), int(row1v2.Y+100), 30, rl.White)
		}
	}

	// draw camera 2X coin images
	rl.BeginMode2D(camera2X)

	if shopshiplr {
		originv2 := rl.NewVector2(120, 0)
		originv2shadow := rl.NewVector2(116, -4)
		shipflamev2 := rl.NewVector2(87, 25)
		if playership != ufo {
			rl.DrawCircle(int(shopshipimage.X+48), int(shopshipimage.Y+60), shopscannerradius, scannercolor)
			rl.DrawTexturePro(imgs, playership, shopshipimage, originv2shadow, -90, rl.Black)
			rl.DrawTexturePro(imgs, playership, shopshipimage, originv2, -90, rl.White)
			rl.DrawTexturePro(imgs, shipflame, shopshipflame, shipflamev2, -90, rl.Fade(rl.White, shipflamefade))

			switch playership {
			case ship2:
				rl.DrawText("- 1 damage taken", int(shopshipimage.X+159), int(shopshipimage.Y+55), 20, rl.Maroon)
				rl.DrawText("- 1 damage taken", int(shopshipimage.X+160), int(shopshipimage.Y+54), 20, rl.White)
			case ship3:
				rl.DrawText("+ 1 attack", int(shopshipimage.X+159), int(shopshipimage.Y+55), 20, rl.Maroon)
				rl.DrawText("+ 1 attack", int(shopshipimage.X+160), int(shopshipimage.Y+54), 20, rl.White)
			case ship4:
				rl.DrawText("+ 20% speed", int(shopshipimage.X+159), int(shopshipimage.Y+55), 20, rl.Maroon)
				rl.DrawText("+ 20% speed", int(shopshipimage.X+160), int(shopshipimage.Y+54), 20, rl.White)
			}

		} else {
			rl.DrawCircle(int(shopshipimage.X+48), int(shopshipimage.Y+60), shopscannerradius, scannercolor)
			rl.DrawTexturePro(imgs, playership, shopshipimageufo, originv2shadow, -90, rl.Black)
			rl.DrawTexturePro(imgs, playership, shopshipimageufo, originv2, -90, rl.White)

			rl.DrawText("+ 33% speed", int(shopshipimageufo.X+119), int(shopshipimageufo.Y+46), 20, rl.Maroon)
			rl.DrawText("+ 33% speed", int(shopshipimageufo.X+120), int(shopshipimageufo.Y+45), 20, rl.White)
			rl.DrawText("- 1 damage taken", int(shopshipimageufo.X+119), int(shopshipimageufo.Y+66), 20, rl.Maroon)
			rl.DrawText("- 1 damage taken", int(shopshipimageufo.X+120), int(shopshipimageufo.Y+65), 20, rl.White)
			rl.DrawText("+ 1 attack", int(shopshipimageufo.X+119), int(shopshipimageufo.Y+86), 20, rl.Maroon)
			rl.DrawText("+ 1 attack", int(shopshipimageufo.X+120), int(shopshipimageufo.Y+85), 20, rl.White)
		}
		if currenttarget == lasermarker5 {
			lasertargetv2 := rl.NewVector2(shopshipimage.X-150, shopshipimage.Y+30)
			rl.DrawTextureRec(imgs, lasermarker590, lasertargetv2, rl.White)
		} else {
			lasertargetv2 := rl.NewVector2(shopshipimage.X-150, shopshipimage.Y+42)
			rl.DrawTextureRec(imgs, currenttarget, lasertargetv2, rl.White)
		}
	} else {
		originv2 := rl.NewVector2(0, 0)
		originv2shadow := rl.NewVector2(-4, -4)
		shipflamev2 := rl.NewVector2(-63, -175)
		if playership != ufo {
			rl.DrawCircle(int(shopshipimage.X-44), int(shopshipimage.Y+60), shopscannerradius, scannercolor)
			rl.DrawTexturePro(imgs, playership, shopshipimage, originv2shadow, 90, rl.Black)
			rl.DrawTexturePro(imgs, playership, shopshipimage, originv2, 90, rl.White)
			rl.DrawTexturePro(imgs, shipflame, shopshipflame, shipflamev2, 90, rl.Fade(rl.White, shipflamefade))

			switch playership {
			case ship2:
				rl.DrawText("- 1 damage taken", int(shopshipimage.X-309), int(shopshipimage.Y+46), 20, rl.Maroon)
				rl.DrawText("- 1 damage taken", int(shopshipimage.X-310), int(shopshipimage.Y+45), 20, rl.White)
			case ship3:
				rl.DrawText("+ 1 attack", int(shopshipimage.X-249), int(shopshipimage.Y+46), 20, rl.Maroon)
				rl.DrawText("+ 1 attack", int(shopshipimage.X-250), int(shopshipimage.Y+45), 20, rl.White)
			case ship4:
				rl.DrawText("+ 20% speed", int(shopshipimage.X-269), int(shopshipimage.Y+46), 20, rl.Maroon)
				rl.DrawText("+ 20% speed", int(shopshipimage.X-270), int(shopshipimage.Y+45), 20, rl.White)
			}
		} else {
			originv2 = rl.NewVector2(-30, 0)
			originv2shadow = rl.NewVector2(-34, -4)
			rl.DrawCircle(int(shopshipimage.X-45), int(shopshipimage.Y+60), shopscannerradius, scannercolor)
			rl.DrawTexturePro(imgs, playership, shopshipimageufo, originv2shadow, 90, rl.Black)
			rl.DrawTexturePro(imgs, playership, shopshipimageufo, originv2, 90, rl.White)

			rl.DrawText("+ 33% speed", int(shopshipimageufo.X-289), int(shopshipimageufo.Y+46), 20, rl.Maroon)
			rl.DrawText("+ 33% speed", int(shopshipimageufo.X-290), int(shopshipimageufo.Y+45), 20, rl.White)
			rl.DrawText("- 1 damage taken", int(shopshipimageufo.X-289), int(shopshipimageufo.Y+66), 20, rl.Maroon)
			rl.DrawText("- 1 damage taken", int(shopshipimageufo.X-290), int(shopshipimageufo.Y+65), 20, rl.White)
			rl.DrawText("+ 1 attack", int(shopshipimageufo.X-289), int(shopshipimageufo.Y+86), 20, rl.Maroon)
			rl.DrawText("+ 1 attack", int(shopshipimageufo.X-290), int(shopshipimageufo.Y+85), 20, rl.White)
		}
		if currenttarget == lasermarker5 {
			lasertargetv2 := rl.NewVector2(shopshipimage.X+150, shopshipimage.Y+30)
			rl.DrawTextureRec(imgs, lasermarker590, lasertargetv2, rl.White)
		} else {
			lasertargetv2 := rl.NewVector2(shopshipimage.X+150, shopshipimage.Y+42)
			rl.DrawTextureRec(imgs, currenttarget, lasertargetv2, rl.White)
		}
	}
	if shopshiplr {
		shopshipimageufo.X -= 6
		shopshipimage.X -= 6
		shopshipflame.X -= 6
		if shopshipimage.X < -75 {
			shopshiplr = false
		}
	} else {
		shopshipimageufo.X += 6
		shopshipimage.X += 6
		shopshipflame.X += 6
		if shopshipimage.X > float32(monw/2)+75 {
			shopshiplr = true
		}
	}

	/*
			shipstorev2 := rl.NewVector2(100, float32((monh/2)-200))
			shipstorev2shadow := rl.NewVector2(98, float32((monh/2)-197))
			currentstoreship := playership
			if storeselect == 1 {
				currentstoreship = ship2
			} else if storeselect == 2 {
				currentstoreship = ship3
			} else if storeselect == 3 {
				currentstoreship = ship4
			} else if storeselect == 4 {
				currentstoreship = ufo
			} else {
				currentstoreship = playership
			}


		if storeselect == 9 {
			rl.DrawCircle(int(shipstorev2.X+56), int(shipstorev2.Y+42), shopscannerradius, rl.Fade(rl.DarkGreen, 0.4))
		} else if storeselect == 10 {
			rl.DrawCircle(int(shipstorev2.X+56), int(shipstorev2.Y+42), shopscannerradius, rl.Fade(rl.Yellow, 0.4))
		} else if storeselect == 11 {
			rl.DrawCircle(int(shipstorev2.X+56), int(shipstorev2.Y+42), shopscannerradius, rl.Fade(rl.Red, 0.4))
		} else if storeselect == 12 {
			rl.DrawCircle(int(shipstorev2.X+56), int(shipstorev2.Y+42), shopscannerradius, rl.Fade(rl.Purple, 0.4))
		} else {
			rl.DrawCircle(int(shipstorev2.X+56), int(shipstorev2.Y+42), shopscannerradius, rl.Fade(scannercolor, 0.4))
		}

		rl.DrawTextureRec(imgs, currentstoreship, shipstorev2shadow, rl.Black)
		rl.DrawTextureRec(imgs, currentstoreship, shipstorev2, rl.White)

		lasermarkerv2 := rl.NewVector2((shipstorev2.X*2)+92, (shipstorev2.Y-30)*2)
		rl.DrawTextureRec(imgs, lasermarker1, lasermarkerv2, rl.White)
	*/

	coinv2 := rl.NewVector2(72, 168)
	if !greenshippurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 85
	if !yellowshippurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 85
	if !redshippurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 78
	if !ufopurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 72
	if !target1purchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 70
	if !target2purchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 70
	if !target3purchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 70
	if !target4purchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2 = rl.NewVector2(72, 269)
	if !greenscannerpurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 85
	if !yellowscannerpurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 85
	if !redscannerpurchashed {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 78
	if !purplescannerpurchased {
		rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	}
	coinv2.X += 72
	rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	coinv2.X += 70
	rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	coinv2.X += 70
	rl.DrawTextureRec(imgs, coin, coinv2, rl.White)
	coinv2.X += 70
	rl.DrawTextureRec(imgs, coin, coinv2, rl.White)

	rl.EndMode2D()

	// animate ship flame

	// animate ship scanner
	if shopscannerradiuson {
		shopscannerradius += 3
		shopscannerradius2 += 2

		if shopscannerradius > 60 {
			shopscannerradiuson = false
		}
	} else {
		shopscannerradius -= 3
		shopscannerradius2 -= 2
		if shopscannerradius < 10 {
			shopscannerradiuson = true
		}
	}

	// navigation keys
	if rl.IsKeyPressed(rl.KeySpace) {
		switch storeselect {
		case 0:
			if spacebaroff == false {
				shopon = false
				spacebaroff = true
				framecount = 0
				nextlevelselect = 1
			}
		case 1:
			if greenshippurchased {
				if playership != ship2 {
					playership = ship2
				} else {
					playership = ship1
				}
			} else {
				if coinnumber-5 >= 0 {
					greenshippurchased = true
					playership = ship2
					coinnumber -= 5
				}
			}
		case 2:
			if yellowshippurchased {
				if playership != ship3 {
					playership = ship3
				} else {
					playership = ship1
				}
			} else {
				if coinnumber-5 >= 0 {
					yellowshippurchased = true
					playership = ship3
					coinnumber -= 5
				}
			}
		case 3:
			if redshippurchased {
				if playership != ship4 {
					playership = ship4
				} else {
					playership = ship1
				}
			} else {
				if coinnumber-5 >= 0 {
					redshippurchased = true
					playership = ship4
					coinnumber -= 5
				}
			}
		case 4:
			if ufopurchased {
				if playership != ufo {
					playership = ufo
				} else {
					playership = ship1
				}
			} else {
				if coinnumber-50 >= 0 {
					ufopurchased = true
					playership = ufo
					coinnumber -= 50
				}
			}
		case 5:
			if target1purchased {
				if currenttarget != lasermarker5 {
					currenttarget = lasermarker5
				} else {
					currenttarget = lasermarker1
				}
			} else {
				if coinnumber-5 >= 0 {
					target1purchased = true
					currenttarget = lasermarker5
					coinnumber -= 5
				}
			}
		case 6:
			if target2purchased {
				if currenttarget != lasermarker2 {
					currenttarget = lasermarker2
				} else {
					currenttarget = lasermarker1
				}
			} else {
				if coinnumber-5 >= 0 {
					target2purchased = true
					currenttarget = lasermarker2
					coinnumber -= 5
				}
			}
		case 7:
			if target3purchased {
				if currenttarget != lasermarker3 {
					currenttarget = lasermarker3
				} else {
					currenttarget = lasermarker1
				}
			} else {
				if coinnumber-5 >= 0 {
					target3purchased = true
					currenttarget = lasermarker3
					coinnumber -= 5
				}
			}
		case 8:
			if target4purchased {
				if currenttarget != lasermarker4 {
					currenttarget = lasermarker4
				} else {
					currenttarget = lasermarker1
				}
			} else {
				if coinnumber-5 >= 0 {
					target4purchased = true
					currenttarget = lasermarker4
					coinnumber -= 5
				}
			}
		case 9:
			if greenscannerpurchased {
				if scannercolor != rl.DarkGreen {
					scannercolor = rl.DarkGreen
				} else {
					scannercolor = rl.Blue
				}
			} else {
				if coinnumber-5 >= 0 {
					greenscannerpurchased = true
					scannercolor = rl.DarkGreen
					coinnumber -= 5
				}
			}
		case 10:
			if yellowscannerpurchased {
				if scannercolor != rl.Yellow {
					scannercolor = rl.Yellow
				} else {
					scannercolor = rl.Blue
				}
			} else {
				if coinnumber-5 >= 0 {
					yellowscannerpurchased = true
					scannercolor = rl.Yellow
					coinnumber -= 5
				}
			}
		case 11:
			if redscannerpurchashed {
				if scannercolor != rl.Red {
					scannercolor = rl.Red
				} else {
					scannercolor = rl.Blue
				}
			} else {
				if coinnumber-5 >= 0 {
					redscannerpurchashed = true
					scannercolor = rl.Red
					coinnumber -= 5
				}
			}
		case 12:
			if purplescannerpurchased {
				if scannercolor != rl.Purple {
					scannercolor = rl.Purple
				} else {
					scannercolor = rl.Blue
				}
			} else {
				if coinnumber-5 >= 0 {
					purplescannerpurchased = true
					scannercolor = rl.Purple
					coinnumber -= 5
				}
			}

		}

	}

	if rl.IsKeyPressed(rl.KeyRight) {
		storeselect++
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		storeselect--
		if storeselect < 1 {
			storeselect = 1
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		if storeselect == 0 {
			storeselect = 1
		} else if storeselect == 1 {
			storeselect = 9
		} else if storeselect == 2 {
			storeselect = 10
		} else if storeselect == 3 {
			storeselect = 11
		} else if storeselect == 4 {
			storeselect = 12
		} else if storeselect == 5 {
			storeselect = 13
		} else if storeselect == 6 {
			storeselect = 14
		} else if storeselect == 7 {
			storeselect = 15
		} else if storeselect == 8 {
			storeselect = 16
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		if storeselect == 1 || storeselect == 2 || storeselect == 3 || storeselect == 4 || storeselect == 5 || storeselect == 6 || storeselect == 7 || storeselect == 8 {
			storeselect = 0
		} else if storeselect == 9 {
			storeselect = 1
		} else if storeselect == 10 {
			storeselect = 2
		} else if storeselect == 11 {
			storeselect = 3
		} else if storeselect == 12 {
			storeselect = 4
		} else if storeselect == 13 {
			storeselect = 5
		} else if storeselect == 14 {
			storeselect = 6
		} else if storeselect == 15 {
			storeselect = 7
		} else if storeselect == 16 {
			storeselect = 8
		}

	}

	// select border fade
	if storeselectfadeon {
		storeselectfade += 0.05
		if storeselectfade > 1.0 {
			storeselectfadeon = false
		}
	} else {
		storeselectfade -= 0.05
		if storeselectfade < 0.2 {
			storeselectfadeon = true
		}
	}

}
func drawcredits() { // MARK: drawcredits
	altcloudson = true
	rl.DrawRectangle(0, 0, monw, monh, rl.Black)
	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(creditscolor, creditsfade))
	if framecount%3 == 0 {
		for a := 0; a < 20; a++ {
			rl.DrawPixel(rInt(50, monw-50), rInt(50, monh-50), rl.Black)
			rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 2, 2, rl.Black)
		}
	}

	if creditsfadeon {
		creditsfade -= 0.04
		if creditsfade <= 0.0 {
			randomcolors()
			creditscolor = randomcolorslice[0]
			creditsfadeon = false
		}
	} else {
		creditsfade += 0.04
		if creditsfade >= 1.0 {
			randomcolors()
			creditscolor = randomcolorslice[0]
			creditsfadeon = true
		}
	}

	rl.DrawText("these links all helped make this game possible:", 100, 150, 40, rl.White)

	rl.DrawText("raylib.com", 150, 250, 40, rl.White)
	rl.DrawText("golang.org", 150, 300, 40, rl.White)
	rl.DrawText("opengameart.org", 150, 350, 40, rl.White)
	rl.DrawText("itch.io", 150, 400, 40, rl.White)
	rl.DrawText("kenney.nl", 150, 450, 40, rl.White)
	rl.DrawText("planets - opengameart.org/content/20-planet-sprites", 150, 500, 40, rl.White)

}
func drawoptionsmenu() { // MARK: drawoptionsmenu
	rl.DrawRectangleLines(monw/2-500, monh/2-200, 1000, 400, rl.Black)
	rl.DrawRectangleLines(monw/2-501, monh/2-201, 1002, 402, rl.Black)
	rl.DrawRectangle(monw/2-500, monh/2-200, 1000, 400, rl.Fade(rl.Black, 0.8))

	switch optionsmenuselect {
	case 1:
		rl.DrawRectangle(monw/2-490, monh/2-186, 440, 54, rl.Black)

	}

	texty := monh/2 - 180
	textx := monw/2 - 470
	textx2 := monw/2 + 30

	rl.DrawText("scanlines", textx-2, texty+2, 40, rl.Maroon)
	rl.DrawText("scanlines", textx-1, texty+1, 40, rl.Black)
	rl.DrawText("scanlines", textx, texty, 40, rl.White)

	rl.DrawText("pixel noise", textx-2, texty+52, 40, rl.Maroon)
	rl.DrawText("pixel noise", textx-1, texty+51, 40, rl.Black)
	rl.DrawText("pixel noise", textx, texty+50, 40, rl.White)

	rl.DrawText("clouds", textx-2, texty+102, 40, rl.Maroon)
	rl.DrawText("clouds", textx-1, texty+101, 40, rl.Black)
	rl.DrawText("clouds", textx, texty+100, 40, rl.White)

	rl.DrawText("weather", textx-2, texty+152, 40, rl.Maroon)
	rl.DrawText("weather", textx-1, texty+151, 40, rl.Black)
	rl.DrawText("weather", textx, texty+150, 40, rl.White)

	checkboxrec := rl.NewRectangle(float32(monw/2-100), float32(texty+8), 24, 24)
	rl.GuiLoadStyle("guistyles.rgs")
	rl.GuiCheckBox(checkboxrec, "", scanlineson)
	checkboxrec.Y += 50
	rl.GuiCheckBox(checkboxrec, "", pixelnoiseon)
	checkboxrec.Y += 50
	rl.GuiCheckBox(checkboxrec, "", cloudson)

	rl.DrawText("sound", textx2-1, texty+1, 40, rl.White)
	rl.DrawText("sound", textx2, texty, 40, rl.Black)

}
func drawstartgamescreen() { // MARK: drawstartgamescreen

	pauseon = true
	startmenuon = true

	rl.DrawRectangle(0, 0, monw, monh, rl.Black) // background rectangle

	if startgeneratecomplete == false {
		for a := 0; a < 3; a++ {
			choose := rInt(1, 8)
			switch choose {
			case 1:
				if a == 0 {
					startplanet1 = planet1
				} else if a == 1 {
					startplanet2 = planet1
				} else if a == 2 {
					startplanet3 = planet1
				}
			case 2:
				if a == 0 {
					startplanet1 = planet2
				} else if a == 1 {
					startplanet2 = planet2
				} else if a == 2 {
					startplanet3 = planet2
				}
			case 3:
				if a == 0 {
					startplanet1 = planet3
				} else if a == 1 {
					startplanet2 = planet3
				} else if a == 2 {
					startplanet3 = planet3
				}
			case 4:
				if a == 0 {
					startplanet1 = planet4
				} else if a == 1 {
					startplanet2 = planet4
				} else if a == 2 {
					startplanet3 = planet4
				}
			case 5:
				if a == 0 {
					startplanet1 = planet5
				} else if a == 1 {
					startplanet2 = planet5
				} else if a == 2 {
					startplanet3 = planet5
				}
			case 6:
				if a == 0 {
					startplanet1 = planet6
				} else if a == 1 {
					startplanet2 = planet6
				} else if a == 2 {
					startplanet3 = planet6
				}
			case 7:
				if a == 0 {
					startplanet1 = planet7
				} else if a == 1 {
					startplanet2 = planet7
				} else if a == 2 {
					startplanet3 = planet7
				}
			}
		}

		planety := rFloat32(100, 300)
		for a := 0; a < 3; a++ {
			choose := flipcoin()
			if choose {
				planetx := rFloat32(-500, -300)
				if a == 0 {
					startplanet1v2 = rl.NewVector2(planetx, planety)
				} else if a == 1 {
					startplanet2v2 = rl.NewVector2(planetx, planety)
				} else if a == 2 {
					startplanet3v2 = rl.NewVector2(planetx, planety)
				}
			} else {
				planetx := rFloat32(monw+50, monw+200)
				if a == 0 {
					startplanet1v2 = rl.NewVector2(planetx, planety)
					startplanet1lr = true
				} else if a == 1 {
					startplanet2v2 = rl.NewVector2(planetx, planety)
					startplanet2lr = true
				} else if a == 2 {
					startplanet3v2 = rl.NewVector2(planetx, planety)
					startplanet3lr = true
				}
			}

			planety += rFloat32(300, 450)

		}

		startplanet1xchange = rFloat32(5, 20)
		startplanet2xchange = rFloat32(5, 20)
		startplanet3xchange = rFloat32(5, 20)

		startplanet1ychange = rFloat32(0, 8)
		startplanet2ychange = rFloat32(0, 8)
		startplanet3ychange = rFloat32(0, 8)

		startplayerdirection = rInt(1, 9)

		startgeneratecomplete = true
	}

	rl.DrawTextureRec(imgs, startplanet1, startplanet1v2, rl.Fade(rl.White, 0.7))
	rl.DrawTextureRec(imgs, startplanet2, startplanet2v2, rl.Fade(rl.White, 0.7))
	rl.DrawTextureRec(imgs, startplanet3, startplanet3v2, rl.Fade(rl.White, 0.7))

	if startplanet1lr {
		startplanet1v2.X -= startplanet1xchange
		startplanet1v2.Y += startplanet1ychange
		if startplanet1v2.X < -250 {
			startplanet1lr = false
		}
	} else {
		startplanet1v2.X += startplanet1xchange
		startplanet1v2.Y -= startplanet1ychange
		if startplanet1v2.X > float32(monw+50) {
			startplanet1lr = true
		}
	}
	if startplanet2lr {
		startplanet2v2.X -= startplanet2xchange
		startplanet2v2.Y += startplanet2ychange
		if startplanet2v2.X < -250 {
			startplanet2lr = false
		}
	} else {
		startplanet2v2.X += startplanet2xchange
		startplanet2v2.Y -= startplanet2ychange
		if startplanet2v2.X > float32(monw+50) {
			startplanet2lr = true
		}
	}
	if startplanet3lr {
		startplanet3v2.X -= startplanet3xchange
		startplanet3v2.Y -= startplanet3ychange
		if startplanet3v2.X < -250 {
			startplanet3lr = false
		}
	} else {
		startplanet3v2.X += startplanet3xchange
		startplanet3v2.Y += startplanet3ychange
		if startplanet3v2.X > float32(monw+50) {
			startplanet3lr = true
		}
	}

	if framecount%3 == 0 {
		for a := 0; a < 20; a++ {
			randomcolors()
			pixelcolor1 := randomcolorslice[0]
			rl.DrawPixel(rInt(50, monw-50), rInt(50, monh-50), pixelcolor1)
			reccolor := randomcolorslice[1]
			rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 3, 3, reccolor)
			reccolor2 := randomcolorslice[2]
			rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 2, 2, reccolor2)
		}
	}

	rl.DrawText("Howard Max - Space & Time Detective in...", monw/6-3, monh/3-57, 40, rl.Maroon)
	rl.DrawText("Howard Max - Space & Time Detective in...", monw/6-1, monh/3-59, 40, rl.Black)
	rl.DrawText("Howard Max - Space & Time Detective in...", monw/6, monh/3-60, 40, rl.White)
	rl.DrawText("The Case of the ALIEN CODE", monw/6-3, monh/3+3, 80, rl.DarkGreen)
	rl.DrawText("The Case of the ALIEN CODE", monw/6-1, monh/3+1, 80, rl.Black)
	rl.DrawText("The Case of the ALIEN CODE", monw/6, monh/3, 80, rl.White)
	rl.DrawText("The Case of the ALIEN CODE", monw/6, monh/3, 80, rl.Fade(rl.Green, starttitlefade))

	if startmenuselect > 6 {
		startmenuselect = 1
	}
	if startmenuselect < 1 {
		startmenuselect = 6
	}

	if startmenuselect == 1 {
		rl.DrawRectangle(monw/2-250, monh/2-60, 300, 60, rl.Fade(rl.White, 0.8))
	} else if startmenuselect == 2 {
		rl.DrawRectangle(monw/2-250, monh/2, 300, 60, rl.Fade(rl.White, 0.8))
	} else if startmenuselect == 3 {
		rl.DrawRectangle(monw/2-250, monh/2+60, 300, 60, rl.Fade(rl.White, 0.8))
	} else if startmenuselect == 4 {
		rl.DrawRectangle(monw/2-250, monh/2+120, 300, 60, rl.Fade(rl.White, 0.8))
	} else if startmenuselect == 5 {
		rl.DrawRectangle(monw/2-250, monh/2+180, 300, 60, rl.Fade(rl.White, 0.8))
	} else if startmenuselect == 6 {
		rl.DrawRectangle(monw/2-250, monh/2+240, 300, 60, rl.Fade(rl.White, 0.8))
	}

	rl.DrawText("start", monw/2-202, monh/2-48, 40, rl.DarkGray)
	rl.DrawText("start", monw/2-201, monh/2-49, 40, rl.Black)
	if startmenuselect == 1 {
		rl.DrawText("start", monw/2-200, monh/2-50, 40, rl.Maroon)
	} else {
		rl.DrawText("start", monw/2-200, monh/2-50, 40, rl.White)
	}

	rl.DrawText("options", monw/2-202, monh/2+12, 40, rl.DarkGray)
	rl.DrawText("options", monw/2-201, monh/2+11, 40, rl.Black)
	if startmenuselect == 2 {
		rl.DrawText("options", monw/2-200, monh/2+10, 40, rl.Maroon)
	} else {
		rl.DrawText("options", monw/2-200, monh/2+10, 40, rl.White)
	}

	rl.DrawText("help", monw/2-202, monh/2+73, 40, rl.DarkGray)
	rl.DrawText("help", monw/2-201, monh/2+71, 40, rl.Black)
	if startmenuselect == 3 {
		rl.DrawText("help", monw/2-200, monh/2+70, 40, rl.Maroon)
	} else {
		rl.DrawText("help", monw/2-200, monh/2+70, 40, rl.White)
	}

	rl.DrawText("scores", monw/2-202, monh/2+133, 40, rl.DarkGray)
	rl.DrawText("scores", monw/2-201, monh/2+131, 40, rl.Black)
	if startmenuselect == 4 {
		rl.DrawText("scores", monw/2-200, monh/2+130, 40, rl.Maroon)
	} else {
		rl.DrawText("scores", monw/2-200, monh/2+130, 40, rl.White)
	}

	rl.DrawText("credits", monw/2-202, monh/2+193, 40, rl.DarkGray)
	rl.DrawText("credits", monw/2-201, monh/2+191, 40, rl.Black)
	if startmenuselect == 5 {
		rl.DrawText("credits", monw/2-200, monh/2+190, 40, rl.Maroon)
	} else {
		rl.DrawText("credits", monw/2-200, monh/2+190, 40, rl.White)
	}

	rl.DrawText("exit", monw/2-202, monh/2+253, 40, rl.DarkGray)
	rl.DrawText("exit", monw/2-201, monh/2+251, 40, rl.Black)
	if startmenuselect == 6 {
		rl.DrawText("exit", monw/2-200, monh/2+250, 40, rl.Maroon)
	} else {
		rl.DrawText("exit", monw/2-200, monh/2+250, 40, rl.White)
	}

	switch startplayerdirection {
	case 1:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, -30, rl.White)
		startscreenship.X -= 5
		startscreenship.Y -= 5
		if startscreenship.Y < -50 || startscreenship.X < -50 {
			startplayerdirection = rInt(4, 7)
		}
	case 2:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 0, rl.White)
		startscreenship.Y -= 5
		if startscreenship.Y < -50 || startscreenship.X < -50 || startscreenship.X > float32(monw+50) {
			startplayerdirection = rInt(4, 8)
		}
	case 3:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 30, rl.White)
		startscreenship.X += 5
		startscreenship.Y -= 5
		if startscreenship.Y < -50 || startscreenship.X > float32(monw+50) {
			startplayerdirection = rInt(6, 9)
		}
	case 4:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 90, rl.White)
		startscreenship.X += 5
		if startscreenship.X > float32(monw+50) {
			startplayerdirection = rInt(6, 9)
		}
	case 5:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 120, rl.White)
		startscreenship.X += 5
		startscreenship.Y += 5
		if startscreenship.Y > float32(monh+50) || startscreenship.X > float32(monw+50) {
			choose := rInt(1, 4)
			switch choose {
			case 1:
				startplayerdirection = 1
			case 2:
				startplayerdirection = 2
			case 3:
				startplayerdirection = 8
			}
		}
	case 6:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 180, rl.White)
		startscreenship.Y += 5
		if startscreenship.Y > float32(monh+50) {
			startplayerdirection = rInt(1, 4)
		}
	case 7:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 210, rl.White)
		startscreenship.X -= 5
		startscreenship.Y += 5
		if startscreenship.Y > float32(monh+50) || startscreenship.X < -50 {
			startplayerdirection = rInt(2, 5)
		}
	case 8:
		originv2 := rl.NewVector2(56, 38)
		rl.DrawTexturePro(imgs, playership, startscreenship, originv2, 270, rl.White)
		startscreenship.X -= 5
		if startscreenship.X < -50 {
			startplayerdirection = rInt(3, 6)
		}
	}

	if starttitlefadeon {
		if starttitlefade > 0.0 {
			starttitlefade -= 0.02
		} else if starttitlefade <= 0.0 {
			starttitlefadeon = false
		}
	} else {
		if starttitlefade < 0.8 {
			starttitlefade += 0.02
		} else if starttitlefade >= 0.8 {
			starttitlefadeon = true
		}
	}

	if framecount%60 == 0 {
		startplayerdirection = rInt(1, 9)
	}

}
func getmouseblock() { // MARK: getmouseblock

	for a := 0; a < blocknumber; a++ {
		checkblock := blockmap[a]

		if mousepos.Y > checkblock.xy2.Y && mousepos.Y < checkblock.xy4.Y {
			if mousepos.X > checkblock.xy.X+30 && mousepos.X < checkblock.xy3.X-30 {
				mouseblock = a
			}
		}

	}

}
func update() { // MARK: update
	input()
	playerxy = rl.NewVector2(playerx, playery)
	playercollisionrec = rl.NewRectangle(playerx+66, playery+camera.Target.Y, 112, 75)
	getmouseblock()
	clouds()
	timers()
	animations()
	if newenemyon {
		createenemies()
		newenemyon = false
	}
	movebullets()
	if pauseon == false {
		checkbullets()
		if powerupactive {
			movepowerups()
		}
		playerenemycollisions()
		collectpowerup()
		collectaliencode()
	}
	// draw order list
	drawbullets()
	if pauseon == false {
		drawenemies()
		drawpowerups()
		drawaliencode()
	}
	if bosslevelon {
		drawbosslevel()
	}
	if minibosson {
		drawminiboss()
	}
	drawcoins()

	drawplayer()

	// something
	if coinscreated == false {
		createcoins()
	}

	if switchclouds {
		if altcloudson {
			randomcolors()
			choose := rInt(0, 10)
			cloudcolor = randomcolorslice[choose]
		} else {
			cloudcolor = rl.White
		}
		switchclouds = false
	}

	if endlevel {
		pauseon = true
		interstitialon = true
		nextlevel()
	}
	if creditson {
		drawcredits()
	}
	if shopon {
		drawshop()
	}
	if startlevelon {
		startlevelname()
	}
	if bosslevelon {
		bosslevel()
	}

	if startgamescreenon {
		drawstartgamescreen()
	}

	drawfx()
	if interstitialon == false {
		drawinfobar()
	}
	if optionsmenuon {
		cloudson = false
		pauseon = true
		stoptimer = true
		drawoptionsmenu()
	}

	if debugon {
		debug()
	}
}
func collectaliencode() { // MARK: collectaliencode

	if aliencodeactive {

		checkaliencode := aliencodemap[aliencodeactivenumber]

		aliencoderec := rl.NewRectangle(checkaliencode.v2.X, checkaliencode.v2.Y, 60, 98)
		playercollisionrec4 := rl.NewRectangle(playerx, playery, 112, 75)

		if rl.CheckCollisionRecs(playercollisionrec4, aliencoderec) {

			aliencodemapplayer[aliencodeplayercount] = checkaliencode
			checkaliencode.collected = true
			checkaliencode.displayed = true
			aliencodemap[aliencodeactivenumber] = checkaliencode
			aliencodeactive = false
			aliencodeselected = false
			aliencodeplayercount++
			if aliencodeplayercount == 3 {
				endlevel = true
			}

		}

	}

}
func collectpowerup() { // MARK: collectpowerup

	for a := 0; a < len(powerupsmap); a++ {
		checkpowerup := powerupsmap[a]

		if checkpowerup.active {
			poweruprec := rl.NewRectangle(checkpowerup.v2.X, checkpowerup.v2.Y, 60, 60)
			playercollisionrec3 := rl.NewRectangle(playerx, playery, 112, 75)
			if rl.CheckCollisionRecs(playercollisionrec3, poweruprec) {

				playerpowerups[powerupmapnumber] = checkpowerup
				powerupmapnumber++
				if powerupmapnumber == 5 {
					powerupmapnumber = 0
				}
				checkpowerup.active = false
				powerupsmap[a] = checkpowerup

				if checkpowerup.randommod1 {
					bulletmodifiers.angle += checkpowerup.angle
				} else {
					if bulletmodifiers.angle > 4 {
						bulletmodifiers.angle -= checkpowerup.angle
					} else {
						bulletmodifiers.angle += checkpowerup.angle / 2
					}
				}
				if checkpowerup.randommod2 {
					bulletmodifiers.anglemultiplier += checkpowerup.anglemultiplier
				} else {
					if bulletmodifiers.anglemultiplier > 4 {
						bulletmodifiers.anglemultiplier -= checkpowerup.anglemultiplier
					} else {
						bulletmodifiers.anglemultiplier += checkpowerup.anglemultiplier / 2
					}
				}
				if checkpowerup.randommod3 {
					bulletmodifiers.leftnumber += checkpowerup.leftnumber
				} else {
					if bulletmodifiers.leftnumber > 4 {
						bulletmodifiers.leftnumber -= checkpowerup.leftnumber
					} else {
						bulletmodifiers.leftnumber += checkpowerup.leftnumber / 2
					}
				}
				if checkpowerup.randommod4 {
					bulletmodifiers.rightnumber += checkpowerup.rightnumber
				} else {
					if bulletmodifiers.rightnumber > 4 {
						bulletmodifiers.rightnumber -= checkpowerup.rightnumber
					} else {
						bulletmodifiers.rightnumber += checkpowerup.rightnumber / 2
					}
				}
				if checkpowerup.randommod5 {
					bulletmodifiers.spread += checkpowerup.spread
				} else {
					if bulletmodifiers.spread > 30 {
						bulletmodifiers.spread -= checkpowerup.spread
					} else {
						bulletmodifiers.spread += checkpowerup.spread / 2
					}
				}
				if checkpowerup.randommod6 {
					bulletmodifiers.number += checkpowerup.number
				} else {
					if bulletmodifiers.number > 4 {
						bulletmodifiers.number -= checkpowerup.number
					} else {
						bulletmodifiers.number += checkpowerup.number / 2
					}
				}

				bulletmodifiers.downnumber = checkpowerup.downnumber
				bulletmodifiers.angle2on = checkpowerup.angle2on
				bulletmodifiers.angleon = checkpowerup.angleon
				bulletmodifiers.down = checkpowerup.down
				bulletmodifiers.left = checkpowerup.left
				bulletmodifiers.right = checkpowerup.right
				bulletmodifiers.img = rInt(1, 5)

				bulletmodifierscorrections()

			}
		}

	}

}
func bulletmodifierscorrections() { // MARK: bulletmodifierscorrections

	if bulletmodifiers.rightnumber > 10 {
		bulletmodifiers.rightnumber -= rInt(3, 8)
	}
	if bulletmodifiers.leftnumber > 10 {
		bulletmodifiers.leftnumber -= rInt(3, 8)
	}
	if bulletmodifiers.spread > 100 {
		bulletmodifiers.spread -= rFloat32(20, 80)
	}
	if bulletmodifiers.spread < 20 {
		bulletmodifiers.spread += rFloat32(10, 20)
	}
	if bulletmodifiers.anglemultiplier > 6 {
		bulletmodifiers.anglemultiplier -= rInt(1, 4)
	}
	if bulletmodifiers.angle > 6 {
		bulletmodifiers.angle -= rInt(1, 4)
	}

}
func playerenemycollisions() { // MARK: playerenemycollisions
	for a := 0; a < len(enemymap); a++ {
		checkenemy := enemymap[a]
		if checkenemy.hp != 0 {
			enemyrec := rl.NewRectangle(checkenemy.xy.X, checkenemy.xy.Y, 100, 80)
			playercollisionrec2 := rl.NewRectangle(playerx, playery, 112, 75)
			if rl.CheckCollisionRecs(playercollisionrec2, enemyrec) {
				checkenemy.hp--
				if checkenemy.hp == 0 {
					score += checkenemy.score
					totalscore += checkenemy.score
					explosionenemyx = int(checkenemy.xy.X)
					explosionenemyy = int(checkenemy.xy.Y)
					explosionenemystart = false
					explosionenemy()
					kills++
				}
				enemymap[a] = checkenemy
				collisions++
				collisiontrue = true
				if playerhppause == false {
					if playership == ufo || playership == ship2 {
						playerhp -= 4
					} else {
						playerhp -= 5
					}
					playerhppause = true
					explosionstart = false
					explosion()
				}

			}
		}
	}
}
func movepowerups() { // MARK: movepowerups
	if pauseon == false {
		for a := 0; a < powerupnumber; a++ {

			checkpowerup := powerupsmap[a]

			switch checkpowerup.direction {
			case 1:
				checkpowerup.v2.X += 4
				checkpowerup.v2.Y -= 2
			case 2:
				checkpowerup.v2.X += 4
			case 3:
				checkpowerup.v2.X += 4
				checkpowerup.v2.Y += 2
			case 4:
				checkpowerup.v2.X -= 4
				checkpowerup.v2.Y -= 2
			case 5:
				checkpowerup.v2.X -= 4
			case 6:
				checkpowerup.v2.X -= 4
				checkpowerup.v2.Y += 2
			}

			powerupsmap[a] = checkpowerup

		}
	}
}
func movebullets() { // MARK: movebullets
	if pauseon == false {
		for a := 0; a < len(bulletmap); a++ {
			checkbullet := bulletmap[a]
			if checkbullet.active {
				if checkbullet.left == false && checkbullet.right == false && checkbullet.down == false {
					if checkbullet.xy.Y > 0 {
						checkbullet.xy.Y -= 15
						if bulletmodifiers.angleon {
							if checkbullet.side == 1 {
								checkbullet.xy.X -= checkbullet.xchange
							} else if checkbullet.side == 2 {
								checkbullet.xy.X += checkbullet.xchange
							}
						}
						bulletmap[a] = checkbullet
					} else {
						bulletmap[a] = bullet{}
					}
				} else if checkbullet.left {
					if checkbullet.xy.X > 0 {
						checkbullet.xy.X -= 15
						bulletmap[a] = checkbullet
					} else {
						bulletmap[a] = bullet{}
					}
				} else if checkbullet.right {
					if checkbullet.xy.X < float32(monw) {
						checkbullet.xy.X += 15
						bulletmap[a] = checkbullet
					} else {
						bulletmap[a] = bullet{}
					}
				} else if checkbullet.down {
					if checkbullet.xy.Y < float32(monh) {
						checkbullet.xy.Y += 15
						bulletmap[a] = checkbullet
					} else {
						bulletmap[a] = bullet{}
					}
				}
			}
		}
	}
}
func checkbullets() { // MARK: checkbullets

	for a := 0; a < len(bulletmap); a++ {
		checkbullet := bulletmap[a]
		if checkbullet.active {
			for a := 0; a < len(enemymap); a++ {
				checkenemy := enemymap[a]
				if checkenemy.hp >= 0 {
					enemyrec := rl.NewRectangle(checkenemy.xy.X, checkenemy.xy.Y, 100, 80)
					bulletrec := rl.NewRectangle(checkbullet.xy.X, checkbullet.xy.Y, 20, 20)
					if rl.CheckCollisionRecs(bulletrec, enemyrec) {

						checkenemy.hp--
						if playership == ufo || playership == ship3 {
							checkenemy.hp--
						}

						if checkenemy.hp <= 0 {
							score += checkenemy.score
							totalscore += checkenemy.score
							explosionenemyx = int(checkenemy.xy.X)
							explosionenemyy = int(checkenemy.xy.Y)
							explosionenemystart = false
							explosionenemy()
							kills++
						}
						enemymap[a] = checkenemy
					}
				}
			}
		}
	}

	for a := 0; a < len(minibossbulletmap); a++ {

		checkbullet := minibossbulletmap[a]
		if checkbullet.active {
			bulletrec := rl.NewRectangle(checkbullet.v2.X, checkbullet.v2.Y, 24, 24)
			playercollisionrec := rl.NewRectangle(playerx, playery, 112, 75)

			if rl.CheckCollisionRecs(playercollisionrec, bulletrec) {
				playerhp -= checkbullet.hp
				playerhppause = true
				checkbullet.active = false
				minibossbulletmap[a] = checkbullet
			}

		}
	}

}
func createcoins() { // MARK: createcoins
	time := 0
	for a := 0; a < len(coinsmap); a++ {
		newcoin := coinstruct{}
		newcoin.xy.X = rFloat32((monw/5)/4, ((monw/5)*4)/4)
		newcoin.xy.Y = rFloat32(-40, -25)
		newcoin.time = rInt(2, 6)
		newcoin.rec = rl.NewRectangle(newcoin.xy.X, newcoin.xy.Y, 64, 64)
		timeholder := newcoin.time
		newcoin.time += time
		time += timeholder
		coinsmap[a] = newcoin
	}
	coinscreated = true
}
func createbullets() { // MARK: createbullets
	if bulletmodifiers.number > 1 {

		startxy := playerxy
		startxy.X += 56
		bulletxycheck = startxy

		if int(bulletmodifiers.number)%2 == 0 {
			startxy.X -= bulletmodifiers.spread / 2
			startxy.X -= ((bulletmodifiers.number / 2) - 1) * bulletmodifiers.spread

			halfbullets := (bulletmodifiers.number / 2) - 1
			angletotal := halfbullets * float32(bulletmodifiers.anglemultiplier*bulletmodifiers.angle)

			for a := 0; a < int(bulletmodifiers.number); a++ {
				newbullet := bullet{}
				newbullet.xy = startxy
				newbullet.damage = 2
				newbullet.active = true

				if newbullet.xy.X < playerxy.X-10 {
					newbullet.side = 1
				} else if newbullet.xy.X > playerxy.X+10 {
					newbullet.side = 2
				}

				if bulletmodifiers.angleon {
					newbullet.xchange = float32(bulletmodifiers.anglemultiplier * bulletmodifiers.angle)
				}
				if bulletmodifiers.angle2on {
					if newbullet.side == 1 {
						newbullet.xchange += angletotal
						angletotal -= float32(bulletmodifiers.anglemultiplier * bulletmodifiers.angle)
					} else if newbullet.side == 2 {
						newbullet.xchange += float32(bulletmodifiers.anglemultiplier*bulletmodifiers.angle) + angletotal
						angletotal += float32(bulletmodifiers.anglemultiplier * bulletmodifiers.angle)
					}
				}

				bulletmap[bulletcount] = newbullet

				startxy.X += bulletmodifiers.spread
				bulletcount++
				if bulletcount >= 498 {
					bulletcount = 0
				}
			}
		} else {

			halfbullets := (bulletmodifiers.number / 2) - 1
			angletotal := halfbullets * float32(bulletmodifiers.anglemultiplier*bulletmodifiers.angle)

			numberback := bulletmodifiers.number
			numberback--
			numberback = numberback / 2

			startxy.X -= numberback * bulletmodifiers.spread

			for a := 0; a < int(bulletmodifiers.number); a++ {
				newbullet := bullet{}
				newbullet.xy = startxy
				newbullet.damage = 2
				newbullet.active = true

				if newbullet.xy.X < playerxy.X+46 {
					newbullet.side = 1
				} else if newbullet.xy.X > playerxy.X+56 {
					newbullet.side = 2
				}

				if bulletmodifiers.angleon {
					newbullet.xchange = float32(bulletmodifiers.anglemultiplier * bulletmodifiers.angle)
				}
				if bulletmodifiers.angle2on {
					if newbullet.side == 1 {
						newbullet.xchange += angletotal
						angletotal -= float32(bulletmodifiers.anglemultiplier * bulletmodifiers.angle)
					} else if newbullet.side == 2 {
						newbullet.xchange += float32(bulletmodifiers.anglemultiplier*bulletmodifiers.angle) + angletotal
						angletotal += float32(bulletmodifiers.anglemultiplier * bulletmodifiers.angle)
					}
				}

				bulletmap[bulletcount] = newbullet

				startxy.X += bulletmodifiers.spread

				bulletcount++
				if bulletcount >= 498 {
					bulletcount = 0
				}
			}
		}

		if bulletmodifiers.left {
			startxy = playerxy
			startxy.X -= 40
			startxy.Y += 100
			for a := 0; a < bulletmodifiers.leftnumber; a++ {
				newbullet := bullet{}
				newbullet.xy = startxy
				newbullet.damage = 2
				newbullet.active = true
				newbullet.left = true
				bulletmap[bulletcount] = newbullet

				startxy.Y -= bulletmodifiers.spread

				bulletcount++
				if bulletcount >= 498 {
					bulletcount = 0
				}
			}
		}

		if bulletmodifiers.right {
			startxy = playerxy
			startxy.X += 152
			startxy.Y += 100
			for a := 0; a < bulletmodifiers.rightnumber; a++ {
				newbullet := bullet{}
				newbullet.xy = startxy
				newbullet.damage = 2
				newbullet.active = true
				newbullet.right = true
				bulletmap[bulletcount] = newbullet

				startxy.Y -= bulletmodifiers.spread

				bulletcount++
				if bulletcount >= 498 {
					bulletcount = 0
				}
			}
		}
		if bulletmodifiers.down {

			if bulletmodifiers.downnumber == 1 {
				startxy = playerxy
				startxy.X += 56
				startxy.Y += 50

				newbullet := bullet{}
				newbullet.xy = startxy
				newbullet.damage = 2
				newbullet.active = true
				newbullet.down = true
				bulletmap[bulletcount] = newbullet

				startxy.Y -= bulletmodifiers.spread

				bulletcount++
				if bulletcount >= 498 {
					bulletcount = 0
				}
			}

		}

	} else {
		// center bullet
		newbullet := bullet{}
		newbullet.xy = playerxy
		newbullet.xy.X += 48
		newbullet.damage = 2
		newbullet.active = true

		bulletmap[bulletcount] = newbullet
		bulletcount++
		if bulletcount >= 498 {
			bulletcount = 0
		}
	}
}
func createweather() { // MARK: createrain

	for a := 0; a < len(rainmap); a++ {
		rainx := rFloat32(10, monw-10)
		rainy := rFloat32(-20, monh-20)
		rainmap[a] = rl.NewVector2(rainx, rainy)
	}

	for a := 0; a < len(snowtypemap); a++ {
		snowx := rFloat32(10, monw-10)
		snowy := rFloat32(-20, monh-20)
		snowv2map[a] = rl.NewVector2(snowx, snowy)
		snowtypemap[a] = rInt(1, 5)
	}

}
func createaliencode() { // MARK: createaliencode

	for a := 0; a < len(aliencodemap); a++ {
		aliencodemap[a] = aliencode{}
	}

	count := 0
	for {
		choose := rInt(0, 10)
		checkaliencode := aliencodemap[choose]
		if checkaliencode.created == false {
			switch choose {
			case 1:
				checkaliencode.img = code1
			case 2:
				checkaliencode.img = code2
			case 3:
				checkaliencode.img = code3
			case 4:
				checkaliencode.img = code4
			case 5:
				checkaliencode.img = code5
			case 6:
				checkaliencode.img = code6
			case 7:
				checkaliencode.img = code7
			case 8:
				checkaliencode.img = code8
			case 9:
				checkaliencode.img = code9
			case 10:
				checkaliencode.img = code10
			}
			checkaliencode.created = true

			checkaliencode.v2.X = rFloat32((monw/2)-400, (monw/2)+400)
			checkaliencode.v2.Y = -200

			aliencodemap[choose] = checkaliencode

			count++
			if count == 10 {
				break
			}

		}

	}

}
func createenemies() { // MARK: createenemies

	for a := 0; a < len(enemymap); a++ {
		enemymap[a] = enemy{}
	}

	chooseenemytype := rInt(1, 10)
	switch chooseenemytype {
	case 1:
		enemytype = "v"
	case 2:
		enemytype = "^"
	case 3:
		enemytype = "diagonal2"
	case 4:
		enemytype = "diagonal"
	case 5:
		enemytype = "vdown"
	case 6:
		enemytype = "2xhorizontalflat"
	case 7:
		enemytype = "horizontalflat"
	case 8:
		enemytype = "2xhorizontal"
	case 9:
		enemytype = "horizontal"
	}

	number := rInt(7, 13)
	driftswitch := flipcoin()
	driftamount := rInt(0, 5)

	imgswitch := flipcoin()
	imgrec := enemy1
	imgchoose := rInt(0, 9)
	switch imgchoose {
	case 0:
		imgrec = enemy1
	case 1:
		imgrec = enemy2
	case 2:
		imgrec = enemy3
	case 3:
		imgrec = enemy4
	case 4:
		imgrec = enemy5
	case 5:
		imgrec = enemy6
	case 6:
		imgrec = enemy7
	case 7:
		imgrec = enemy8
	case 8:
		imgrec = enemy9
	}

	switch enemytype {
	case "v":
		xchange := rFloat32(monw/20, monw/10)
		ychange := rFloat32(monw/20, monw/10)
		enemyy := rFloat32(-1500, -1300)
		enemyx := rFloat32(10, 100)
		for a := 0; a < number/2; a++ {
			enemyx += xchange
			enemyy += ychange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			newenemy.active = true
			enemymap[a] = newenemy
		}
		for a := number / 2; a < number; a++ {
			enemyx += xchange
			enemyy -= ychange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			newenemy.active = true
			enemymap[a] = newenemy
		}
	case "^":
		xchange := rFloat32(monw/20, monw/10)
		ychange := rFloat32(monw/20, monw/10)
		enemyy := rFloat32(-200, -100)
		enemyx := rFloat32(10, 100)
		for a := 0; a < number/2; a++ {
			enemyx += xchange
			enemyy -= ychange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			newenemy.active = true
			enemymap[a] = newenemy
		}
		for a := number / 2; a < number; a++ {
			enemyx += xchange
			enemyy += ychange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			newenemy.active = true
			enemymap[a] = newenemy
		}
	case "diagonal2":
		xchange := rFloat32(monw/20, monw/10)
		ychange := rFloat32(monw/20, monw/10)
		enemyy := rFloat32(-200, -100)
		enemyx := rFloat32(10, 100)
		for a := 0; a < number; a++ {
			enemyx += xchange
			enemyy -= ychange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			newenemy.active = true
			enemymap[a] = newenemy
		}
	case "diagonal":
		xchange := rFloat32(monw/20, monw/10)
		ychange := rFloat32(monw/20, monw/10)
		enemyy := rFloat32(-2200, -1800)
		enemyx := rFloat32(10, 100)
		for a := 0; a < number; a++ {
			enemyx += xchange
			enemyy += ychange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
		}
	case "vdown":
		enemyx := float32(monw / 2)
		enemyy := rFloat32(-200, -100)
		enemyx += rFloat32(-100, 100)
		xholder := enemyx
		yholder := enemyy
		xchange := rFloat32(monw/20, monw/10)
		ychange := rFloat32(monw/20, monw/10)
		countholder := 0
		for a := 0; a < number/2; a++ {
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
			enemyx -= xchange
			enemyy -= ychange
			countholder = a
		}
		enemyx = xholder
		enemyy = yholder
		for a := countholder; a < countholder+(number/2); a++ {
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
			enemyx += xchange
			enemyy -= ychange
		}
	case "2xhorizontalflat":
		xchange := rFloat32(monw/20, monw/10)
		enemyy := rFloat32(-300, -200)
		enemyx := rFloat32(10, 100)
		countholder := 0
		for a := 0; a < number; a++ {
			enemyx += xchange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
			countholder = a
		}
		xchange = rFloat32(monw/20, monw/10)
		enemyy -= rFloat32(-400, -100)
		enemyx = rFloat32(10, 100)
		for a := countholder; a < countholder+(number); a++ {
			enemyx += xchange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
		}
	case "horizontalflat":
		xchange := rFloat32(monw/20, monw/10)
		enemyy := rFloat32(-300, -200)
		enemyx := rFloat32(10, 100)
		for a := 0; a < number; a++ {
			enemyx += xchange
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
		}
	case "2xhorizontal":
		xchange := float32(0)
		countholder := 0
		for a := 0; a < number; a++ {
			xchange += rFloat32(100, monw/12)
			enemyx := rFloat32(10, 100)
			enemyx += xchange
			enemyy := rFloat32(-400, -200)
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
			countholder = a
		}
		xchange = float32(0)
		for a := countholder; a < countholder+(number); a++ {
			xchange += rFloat32(100, monw/12)
			enemyx := rFloat32(10, 100)
			enemyx += xchange
			enemyy := rFloat32(-700, -500)
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
		}
	case "horizontal":
		xchange := float32(0)
		for a := 0; a < number; a++ {
			xchange += rFloat32(100, monw/12)
			enemyx := rFloat32(10, 100)
			enemyx += xchange
			enemyy := rFloat32(-400, -200)
			newenemy := enemy{}
			if imgswitch {
				imgchoose := rInt(0, 9)
				switch imgchoose {
				case 0:
					newenemy.img = enemy1
				case 1:
					newenemy.img = enemy2
				case 2:
					newenemy.img = enemy3
				case 3:
					newenemy.img = enemy4
				case 4:
					newenemy.img = enemy5
				case 5:
					newenemy.img = enemy6
				case 6:
					newenemy.img = enemy7
				case 7:
					newenemy.img = enemy8
				case 8:
					newenemy.img = enemy9
				}
			} else {
				newenemy.img = imgrec
			}
			if driftswitch {
				newenemy.drift = rInt(0, 5)
			} else {
				newenemy.drift = driftamount
			}
			newenemy.xy = rl.NewVector2(enemyx, enemyy)
			newenemy.hp = rInt(1, 5)
			newenemy.score = newenemy.hp * 5
			enemymap[a] = newenemy
		}
	}

}
func poweruptimers() { // MARK: poweruptimers

	poweruptimer1complete = false
	poweruptimer2complete = false
	poweruptimer3complete = false
	poweruptimer4complete = false
	poweruptimer5complete = false
	poweruptimer6complete = false

	poweruptimer1 = rInt(1, 4)
	poweruptimer2 = rInt(12, 16)
	poweruptimer3 = rInt(24, 28)
	poweruptimer4 = rInt(36, 40)
	poweruptimer5 = rInt(40, 44)
	poweruptimer6 = rInt(52, 56)

}
func createpowerup() { // MARK: createpowerup

	for a := 0; a < len(powerupsmap); a++ {
		choose := rInt(1, 11)
		newpowerup := powerup{}
		switch choose {
		case 1:
			newpowerup.name = "shoot1"
			newpowerup.img = powerup1
		case 2:
			newpowerup.name = "shoot2"
			newpowerup.img = powerup2
		case 3:
			newpowerup.name = "shoot3"
			newpowerup.img = powerup3
		case 4:
			newpowerup.name = "shoot4"
			newpowerup.img = powerup4
		case 5:
			newpowerup.name = "shoot5"
			newpowerup.img = powerup5
		case 6:
			newpowerup.name = "shoot6"
			newpowerup.img = powerup6
		case 7:
			newpowerup.name = "shoot7"
			newpowerup.img = powerup7
		case 8:
			newpowerup.name = "shoot8"
			newpowerup.img = powerup8
		case 9:
			newpowerup.name = "shoot9"
			newpowerup.img = powerup9
		case 10:
			newpowerup.name = "shoot10"
			newpowerup.img = powerup10
		}
		newpowerup.side = flipcoin()
		if newpowerup.side {
			newpowerup.direction = rInt(1, 4)
			newpowerup.v2.X = rFloat32(-280, -140)
		} else {
			newpowerup.direction = rInt(4, 7)
			newpowerup.v2.X = rFloat32(monw+140, monw+280)
		}
		newpowerup.v2.Y = rFloat32((monh/2)-300, (monh/2)+300)
		newpowerup.active = true

		newpowerup.angle2on = flipcoin()
		newpowerup.angleon = flipcoin()
		newpowerup.left = flipcoin()
		newpowerup.right = flipcoin()
		newpowerup.down = flipcoin()
		newpowerup.anglemultiplier = rInt(1, 5)
		newpowerup.angle = rInt(1, 5)
		newpowerup.number = rFloat32(1, 4)
		newpowerup.spread = rFloat32(30, 51)
		newpowerup.leftnumber = rInt(1, 5)
		newpowerup.rightnumber = rInt(1, 5)
		newpowerup.downnumber = rInt(0, 2)
		newpowerup.randommod1 = flipcoin()
		newpowerup.randommod2 = flipcoin()
		newpowerup.randommod3 = flipcoin()
		newpowerup.randommod4 = flipcoin()
		newpowerup.randommod5 = flipcoin()
		newpowerup.randommod6 = flipcoin()

		powerupsmap[a] = newpowerup

	}
	powerupnumber = rInt(1, 3)

}
func explosionenemy() { // MARK: explosionenemy

	// enemy explosions
	if explosionenemystart == false {

		for a := 0; a < 10; a++ {
			newexplosion := explosioncircle{}
			newexplosion.x = explosionenemyx + (rInt(-150, 150))
			newexplosion.y = explosionenemyy + (rInt(-150, 150))
			newexplosion.radius = rFloat32(20, 100)
			newexplosion.fade = rFloat32(4, 11)
			newexplosion.fade = (newexplosion.fade / 10)
			choosecolor := rolldice()
			switch choosecolor {
			case 1:
				newexplosion.color = explosioncolor1
			case 2:
				newexplosion.color = explosioncolor2
			case 3:
				newexplosion.color = explosioncolor3
			case 4:
				newexplosion.color = explosioncolor4
			case 5:
				newexplosion.color = explosioncolor5
			case 6:
				newexplosion.color = explosioncolor6
			}
			explosionsenemycircles[a] = newexplosion
		}
		explosionenemyon = true
		explosionenemystart = true
	}

	if explosionenemyon {

		for a := 0; a < 10; a++ {
			checkexplosion := explosionsenemycircles[a]
			rl.DrawCircle(checkexplosion.x, checkexplosion.y, checkexplosion.radius, rl.Fade(checkexplosion.color, checkexplosion.fade))
			rl.DrawCircleLines(checkexplosion.x, checkexplosion.y, checkexplosion.radius, rl.Black)
			rl.DrawCircleLines(checkexplosion.x, checkexplosion.y, checkexplosion.radius+1, rl.Black)
			rl.DrawCircleLines(checkexplosion.x, checkexplosion.y, checkexplosion.radius+2, rl.Black)
			if checkexplosion.radius > 7 {
				checkexplosion.radius -= 5
			}
			checkexplosion.y += 10
			explosionsenemycircles[a] = checkexplosion
		}
	}

}
func explosion() { // MARK: explosion

	// player explosions
	if explosionstart == false {
		explosionx = int(playerx + (playerw / 2))
		explosiony = int(playery + (playerh / 2))

		for a := 0; a < 10; a++ {
			newexplosion := explosioncircle{}
			newexplosion.x = explosionx + (rInt(-150, 150))
			newexplosion.y = explosiony + (rInt(-150, 150))
			newexplosion.radius = rFloat32(20, 100)
			newexplosion.fade = rFloat32(4, 11)
			newexplosion.fade = (newexplosion.fade / 10)
			choosecolor := rolldice()
			switch choosecolor {
			case 1:
				newexplosion.color = explosioncolor1
			case 2:
				newexplosion.color = explosioncolor2
			case 3:
				newexplosion.color = explosioncolor3
			case 4:
				newexplosion.color = explosioncolor4
			case 5:
				newexplosion.color = explosioncolor5
			case 6:
				newexplosion.color = explosioncolor6
			}
			explosionscircles[a] = newexplosion
		}
		explosionon = true
		explosionstart = true
	}

	if explosionon {
		for a := 0; a < 10; a++ {
			checkexplosion := explosionscircles[a]
			rl.DrawCircle(checkexplosion.x, checkexplosion.y, checkexplosion.radius, rl.Fade(checkexplosion.color, checkexplosion.fade))
			rl.DrawCircleLines(checkexplosion.x, checkexplosion.y, checkexplosion.radius, rl.Black)
			rl.DrawCircleLines(checkexplosion.x, checkexplosion.y, checkexplosion.radius+1, rl.Black)
			rl.DrawCircleLines(checkexplosion.x, checkexplosion.y, checkexplosion.radius+2, rl.Black)
			if checkexplosion.radius > 10 {
				checkexplosion.radius -= 5
			}
			checkexplosion.y += 10
			explosionscircles[a] = checkexplosion
		}
	}

}
func animations() { // MARK: animations

	// level tile fade
	if framecount%2 == 0 {
		if whitefadeon {
			whitefade1 -= 0.1
			whitefade2 -= 0.1
			whitefade3 -= 0.1
			whitefade4 -= 0.1
			whitefade5 -= 0.1

			if whitefade5 <= 0.6 {
				whitefadeon = false
			}
		} else {
			whitefade1 += 0.1
			whitefade2 += 0.1
			whitefade3 += 0.1
			whitefade4 += 0.1
			whitefade5 += 0.1

			if whitefade5 >= 1.0 {
				whitefadeon = true
			}
		}

	}

	// coin animations
	if framecount%2 == 0 {
		coin.X += 16
		if coin.X > 110 {
			coin.X = 41
		}
	}
	// bullet animations
	if framecount%2 == 0 {
		bullet1.X += 20
		if bullet1.X > 121 {
			bullet1.X = 41
		}
		bullet2.X += 20
		if bullet2.X > 450 {
			bullet2.X = 362
		}
		bullet3.X += 20
		if bullet3.X > 450 {
			bullet3.X = 362
		}
		bullet4.X += 20
		if bullet4.X > 560 {
			bullet4.X = 473
		}
	}
	// boss lines fade
	if bosslevelon {
		if framecount%3 == 0 {
			if bosslinesfadeon {
				bosslinesfade += 0.1
				if bosslinesfade >= 0.5 {
					bosslinesfadeon = false
				}
			} else {
				bosslinesfade -= 0.1
				if bosslinesfade <= 0.2 {
					bosslinesfadeon = true
				}
			}
		}
	}

	// alien code shadow fade
	if framecount%2 == 0 {
		if aliencodefadeon {
			aliencodefade += 0.1
			if aliencodefade >= 1.0 {
				aliencodefadeon = false
			}
		} else {
			aliencodefade -= 0.1
			if aliencodefade <= 0.6 {
				aliencodefadeon = true
			}
		}
	}

	// ship flame flicker
	choosefade := rolldice()
	switch choosefade {
	case 1:
		shipflamefade = 0.9
	case 2:
		shipflamefade = 0.8
	case 3:
		shipflamefade = 0.7
	case 4:
		shipflamefade = 0.6
	case 5:
		shipflamefade = 0.5
	case 6:
		shipflamefade = 0.4
	}
	// bullet color fade
	if bulletfadeon {
		bulletfade += 0.1
		if bulletfade >= 1.0 {
			bulletfadeon = false
		}
	} else {
		bulletfade -= 0.1
		if bulletfade <= 0.7 {
			bulletfadeon = true
		}
	}
	// animate score color
	if framecount%15 == 0 {
		if scorecolor == rl.Yellow {
			scorecolor = rl.Red
		} else if scorecolor == rl.Red {
			scorecolor = rl.Green
		} else if scorecolor == rl.Green {
			scorecolor = rl.Blue
		} else if scorecolor == rl.Blue {
			scorecolor = rl.Magenta
		} else if scorecolor == rl.Magenta {
			scorecolor = rl.Orange
		} else if scorecolor == rl.Orange {
			scorecolor = rl.Purple
		} else if scorecolor == rl.Purple {
			scorecolor = rl.SkyBlue
		} else if scorecolor == rl.SkyBlue {
			scorecolor = rl.DarkGreen
		} else if scorecolor == rl.DarkGreen {
			scorecolor = rl.Yellow
		}
	}

	// move enemies
	if pauseon == false {
		for a := 0; a < len(enemymap); a++ {
			checkenemy := enemymap[a]
			if checkenemy.hp != 0 {
				checkenemy.xy.Y += 10
				switch checkenemy.drift {
				case 1:
					checkenemy.xy.X -= 2
				case 2:
					checkenemy.xy.X += 2
				case 3:
					checkenemy.xy.X -= 2
					checkenemy.xy.Y += 5
				case 4:
					checkenemy.xy.X += 2
					checkenemy.xy.Y += 5

				}
				enemymap[a] = checkenemy
			}
		}
	}
}
func timers() { // MARK: timers

	// create powerups
	if !poweruptimer1complete {
		if seconds == poweruptimer1 {
			createpowerup()
			powerupactive = true
			poweruptimer1complete = true
		}
	}
	if !poweruptimer2complete {
		if seconds == poweruptimer2 {
			createpowerup()
			powerupactive = true
			poweruptimer2complete = true
		}
	}
	if !poweruptimer3complete {
		if seconds == poweruptimer3 {
			createpowerup()
			powerupactive = true
			poweruptimer3complete = true
		}
	}
	if !poweruptimer4complete {
		if seconds == poweruptimer4 {
			createpowerup()
			powerupactive = true
			poweruptimer4complete = true
		}
	}
	if !poweruptimer5complete {
		if seconds == poweruptimer5 {
			createpowerup()
			powerupactive = true
			poweruptimer5complete = true
		}
	}
	if !poweruptimer6complete {
		if seconds == poweruptimer6 {
			createpowerup()
			powerupactive = true
			poweruptimer6complete = true
		}
	}

	// create new enemies
	if seconds > 1 && framecount%((fps*8)+1) == 0 {
		newenemies = true
	}
	if newenemies && seconds > 1 && framecount%((fps*9)+2) == 0 {
		createenemies()
		newenemies = false
	}

	// update explosions
	if explosionon {
		explosion()
		if framecount%30 == 0 {
			explosiontimer--
			if explosiontimer <= 0 {
				explosionon = false
				explosiontimer = 1
			}
		}
	}
	if explosionenemyon {
		explosionenemy()
		if framecount%30 == 0 {
			explosionenemytimer--
			if explosionenemytimer <= 0 {
				explosionenemyon = false
				explosionenemytimer = 1
			}
		}
	}

	// collision pause
	if playerhppause {
		if framecount%30 == 0 {
			playerhptimer--
			if playerhptimer <= 0 {
				playerhppause = false
				playerhptimer = 1
			}
		}

	}
}
func clouds() { // MARK: clouds

	if cloudscount == 0 {
		if framecount%60 == 0 {
			cloud1lr = flipcoin()
			cloud2lr = flipcoin()
			cloud3lr = flipcoin()
			cloud1drifton = flipcoin()
			cloud2drifton = flipcoin()
			cloud3drifton = flipcoin()
			cloud1drift = rFloat32(1, 6)
			cloud2drift = rFloat32(1, 6)
			cloud3drift = rFloat32(1, 6)
			cloud1speed = rFloat32(8, 25)
			cloud2speed = rFloat32(8, 25)
			cloud3speed = rFloat32(8, 25)
			createclouds = true
			cloudscount = rInt(1, 4)

			zoomlevel := rFloat32(0, 11)
			zoomlevel = zoomlevel / 10
			zoomlevel++
			cameraclouds.Zoom = zoomlevel

		}
	}
	if cloudscount != 0 && createclouds {
		if cloud1lr {
			cloudx := rFloat32(-600, -450)
			cloudy := rFloat32(10, monh-200)
			cloud1v2 = rl.NewVector2(cloudx, cloudy)
			if cloudy > float32(monh/2) {
				cloud1drift = -(cloud1drift)
			}
		} else {
			cloudx := rFloat32(monw+10, monw+100)
			cloudy := rFloat32(10, monh-200)
			cloud1v2 = rl.NewVector2(cloudx, cloudy)
			if cloudy > float32(monh/2) {
				cloud1drift = -(cloud1drift)
			}
		}
		if cloud2lr {
			cloudx := rFloat32(-600, -450)
			cloudy := rFloat32(10, monh-200)
			cloud2v2 = rl.NewVector2(cloudx, cloudy)
			if cloudy > float32(monh/2) {
				cloud1drift = -(cloud1drift)
			}
		} else {
			cloudx := rFloat32(monw+10, monw+100)
			cloudy := rFloat32(10, monh-200)
			cloud2v2 = rl.NewVector2(cloudx, cloudy)
			if cloudy > float32(monh/2) {
				cloud1drift = -(cloud1drift)
			}
		}
		if cloud3lr {
			cloudx := rFloat32(-600, -450)
			cloudy := rFloat32(10, monh-200)
			cloud3v2 = rl.NewVector2(cloudx, cloudy)
			if cloudy > float32(monh/2) {
				cloud1drift = -(cloud1drift)
			}
		} else {
			cloudx := rFloat32(monw+10, monw+100)
			cloudy := rFloat32(10, monh-200)
			cloud3v2 = rl.NewVector2(cloudx, cloudy)
			if cloudy > float32(monh/2) {
				cloud1drift = -(cloud1drift)
			}
		}
		createclouds = false
	}

	if cloudscount != 0 && startclouds {
		if cloudscount == 1 {
			choose := rInt(1, 4)
			if choose == 1 {
				cloud1active = true
				cloud2active = false
				cloud3active = false
			} else if choose == 2 {
				cloud1active = false
				cloud2active = true
				cloud3active = false
			} else if choose == 3 {
				cloud1active = false
				cloud2active = false
				cloud3active = true
			}
		} else if cloudscount == 2 {
			choose := rInt(1, 4)
			choose2 := rInt(1, 4)
			if choose == 1 {
				cloud1active = true
				cloud2active = false
				cloud3active = false
			} else if choose == 2 {
				cloud1active = false
				cloud2active = true
				cloud3active = false
			} else if choose == 3 {
				cloud1active = false
				cloud2active = false
				cloud3active = true
			}
			if choose2 == 1 {
				if cloud1active == false {
					cloud1active = true
				} else {
					if flipcoin() {
						cloud2active = true
					} else {
						cloud3active = true
					}
				}
			} else if choose2 == 2 {
				if cloud2active == false {
					cloud2active = true
				} else {
					if flipcoin() {
						cloud1active = true
					} else {
						cloud3active = true
					}
				}
			} else if choose2 == 3 {
				if cloud3active == false {
					cloud3active = true
				} else {
					if flipcoin() {
						cloud1active = true
					} else {
						cloud2active = true
					}
				}
			}
		} else if cloudscount == 3 {
			cloud1active = true
			cloud2active = true
			cloud3active = true
		}

		startclouds = false
	}

	if cloud1active {
		if cloud1drifton {
			cloud1v2.Y += cloud1drift
		}
		if cloud1lr {
			cloud1v2.X += cloud1speed
			if cloud1v2.X > float32(monw) {
				cloud1active = false
				cloudscount--
			}
		} else {
			cloud1v2.X -= cloud1speed
			if cloud1v2.X < float32(-450) {
				cloud1active = false
				cloudscount--
			}
		}
	}

	if cloud2active {
		if cloud2drifton {
			cloud2v2.Y += cloud2drift
		}
		if cloud2lr {
			cloud2v2.X += cloud2speed
			if cloud2v2.X > float32(monw) {
				cloud2active = false
				cloudscount--
			}
		} else {
			cloud2v2.X -= cloud2speed
			if cloud2v2.X < float32(-450) {
				cloud2active = false
				cloudscount--
			}
		}
	}

	if cloud3active {
		if cloud3drifton {
			cloud3v2.Y += cloud3drift
		}
		if cloud3lr {
			cloud3v2.X += cloud3speed
			if cloud3v2.X > float32(monw) {
				cloud3active = false
				cloudscount--
			}
		} else {
			cloud3v2.X -= cloud3speed
			if cloud3v2.X < float32(-450) {
				cloud3active = false
				cloudscount--
			}
		}
	}

	if cloudscount < 0 {
		cloudscount = 0
	}

	if cloudscount == 0 {
		startclouds = true
	}

}
func input() { // MARK: input
	if rl.IsKeyPressed(rl.KeyF7) {
		if minibosson {
			minibosson = false
		} else {
			minibosson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF6) {
		coinnumber += 100
	}
	if rl.IsKeyPressed(rl.KeyF5) {
		seconds = 59
	}
	if rl.IsKeyPressed(rl.KeyF4) {
		if shopon {
			shopon = false
			pauseon = false
			stoptimer = false
		} else {
			shopon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyF3) {
		if creditson {
			creditson = false
			altcloudson = false
			switchclouds = true
		} else {
			creditson = true
			altcloudson = true
			switchclouds = true
		}
	}

	if rl.IsKeyPressed(rl.KeyF2) {
		if startgamescreenon {
			startgamescreenon = false
			pauseon = false
		} else {
			startgamescreenon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if optionsmenuon {
			optionsmenuon = false
			pauseon = false
		} else {
			optionsmenuon = true
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		if pauseon == false {
			createbullets()
		} else {
			if interstitialon && nextlevelselect == 1 {
				if spacebaroff == false && shopon == false {
					interstitialon = false
					endlevel = false
					stoptimer = true
					startlevelon = true
				}
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyPause) {
		if pauseon {
			pauseon = false
		} else {
			pauseon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp9) {
		pauseon = true
		bosslevelon = true
	}
	if rl.IsKeyPressed(rl.KeyKp8) {
		createaliencode()
		aliencodeactive = true
		aliencodeselected = false
	}
	if rl.IsKeyPressed(rl.KeyKp7) {
		createpowerup()
		powerupactive = true
	}
	if rl.IsKeyPressed(rl.KeyKp6) {
		newenemyon = true
	}
	if rl.IsKeyPressed(rl.KeyKp5) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp4) {
		explosionstart = false
		explosion()
	}
	if rl.IsKeyPressed(rl.KeyKp3) {
		if scanlineson {
			scanlineson = false
		} else {
			scanlineson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp2) {
		if sceneryon {
			sceneryon = false
		} else {
			sceneryon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp1) {
		if cloudson {
			cloudson = false
		} else {
			cloudson = true
		}
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		if playerx > 100 {
			if playership == ship4 {
				playerx -= 72
			} else if playership == ufo {
				playerx -= 80
			} else {
				playerx -= 60
			}
		}
		playerdirection = 1
	} else if rl.IsKeyDown(rl.KeyRight) {
		if playerx < float32(monw-200) {
			if playership == ship4 {
				playerx += 72
			} else if playership == ufo {
				playerx += 80
			} else {
				playerx += 60
			}
		}
		playerdirection = 3
	} else {
		playerdirection = 2
	}
	if rl.IsKeyDown(rl.KeyUp) {
		if playery > 100 {
			if playership == ship4 {
				playery -= 72
			} else if playership == ufo {
				playery -= 80
			} else {
				playery -= 60
			}
		}
	}
	if rl.IsKeyDown(rl.KeyDown) {
		if playery < float32(monh-200) {
			if playership == ship4 {
				playery += 72
			} else if playership == ufo {
				playery += 80
			} else {
				playery += 60
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		if startmenuon {
			startmenuselect--
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		if startmenuon {
			startmenuselect++
		}
	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {
		camera.Zoom += 0.1
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		camera.Zoom -= 0.1
	}
	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}
}
func resetimers() { // MARK: resetlevel
	if resetimerscomplete == false {
		seconds = 0
		framecount = 0
		resetimerscomplete = true
	}
	nextblock = blocknumber - screenblocknumber*2
}
func nextlevel() { // MARK: nextlevel
	if interstitialon && shopon == false {
		rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Black, endlevelfade))

		if framecount%3 == 0 {
			for a := 0; a < 20; a++ {
				randomcolors()
				pixelcolor1 := randomcolorslice[0]
				rl.DrawPixel(rInt(50, monw-50), rInt(50, monh-50), pixelcolor1)
				reccolor := randomcolorslice[1]
				rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 3, 3, reccolor)
				reccolor2 := randomcolorslice[2]
				rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 2, 2, reccolor2)
			}
		}

		if endlevelfade < 1.0 {
			endlevelfade += 0.02
		} else {
			totalcollisions += collisions
			totalkills += kills

			resetimers()
			poweruptimers()
			clearmaps()
			createlevel()
			coinscreated = false
			nighton = flipcoin()
			if rolldice() > 3 {
				if flipcoin() {
					rainon = true
					snowon = false
				} else {
					snowon = true
					rainon = false
				}
			}

			if framecount == 60 {
				spacebaroff = false
			}

			if altcloudson {
				randomcolors()
				choose := rInt(0, 10)
				cloudcolor = randomcolorslice[choose]
			} else {
				cloudcolor = rl.White
			}

			collisionsTEXT := strconv.Itoa(collisions)
			killsTEXT := strconv.Itoa(kills)
			scoreTEXT := strconv.Itoa(score)

			rl.DrawText("level complete", monw/3-3, monh/3+3, 80, rl.Maroon)
			rl.DrawText("level complete", monw/3-1, monh/3+1, 80, rl.Black)
			rl.DrawText("level complete", monw/3, monh/3, 80, rl.White)

			rl.DrawText("collisions", monw/3+38, (monh/3)+122, 40, rl.Yellow)
			rl.DrawText("collisions", monw/3+39, (monh/3)+121, 40, rl.Black)
			rl.DrawText("collisions", monw/3+40, (monh/3)+120, 40, rl.White)

			rl.DrawText("kills", monw/3+38, (monh/3)+182, 40, rl.Yellow)
			rl.DrawText("kills", monw/3+39, (monh/3)+181, 40, rl.Black)
			rl.DrawText("kills", monw/3+40, (monh/3)+180, 40, rl.White)

			rl.DrawText("score", monw/3+38, (monh/3)+242, 40, rl.Yellow)
			rl.DrawText("score", monw/3+39, (monh/3)+241, 40, rl.Black)
			rl.DrawText("score", monw/3+40, (monh/3)+240, 40, rl.White)

			rl.DrawText(collisionsTEXT, monw/2+100-2, (monh/3)+122, 40, totalscolor)
			rl.DrawText(collisionsTEXT, monw/2+100-1, (monh/3)+121, 40, rl.Black)
			rl.DrawText(collisionsTEXT, monw/2+100, (monh/3)+120, 40, rl.White)

			rl.DrawText(killsTEXT, monw/2+100-2, (monh/3)+182, 40, totalscolor)
			rl.DrawText(killsTEXT, monw/2+100-1, (monh/3)+181, 40, rl.Black)
			rl.DrawText(killsTEXT, monw/2+100, (monh/3)+180, 40, rl.White)

			rl.DrawText(scoreTEXT, monw/2+100-2, (monh/3)+242, 40, totalscolor)
			rl.DrawText(scoreTEXT, monw/2+100-1, (monh/3)+241, 40, rl.Black)
			rl.DrawText(scoreTEXT, monw/2+100, (monh/3)+240, 40, rl.White)

			if nextlevelselect == 0 {
				rl.DrawRectangle(monw/2-304, monh/3+346, 208, 68, selectboxborder)
			} else {
				rl.DrawRectangle(monw/2-4, monh/3+346, 208, 68, selectboxborder)
			}

			if framecount%10 == 0 {
				if selectboxborder == rl.Green {
					selectboxborder = rl.DarkGreen
				} else {
					selectboxborder = rl.Green
				}
			}

			if rl.IsKeyPressed(rl.KeyRight) {
				nextlevelselect++
				if nextlevelselect > 1 {
					nextlevelselect = 0
				}
			} else if rl.IsKeyPressed(rl.KeyLeft) {
				nextlevelselect--
				if nextlevelselect < 0 {
					nextlevelselect = 1
				}
			}

			if rl.IsKeyPressed(rl.KeySpace) && nextlevelselect == 0 {
				shopon = true
			}

			rl.DrawRectangle(monw/2, monh/3+350, 200, 60, rl.Maroon)
			rl.DrawText("next", monw/2+53, monh/3+362, 40, rl.Black)
			rl.DrawText("next", monw/2+55, monh/3+360, 40, rl.White)

			rl.DrawRectangle(monw/2-300, monh/3+350, 200, 60, rl.Maroon)
			rl.DrawText("shop", monw/2-252, monh/3+362, 40, rl.Black)
			rl.DrawText("shop", monw/2-250, monh/3+360, 40, rl.White)

			if framecount%6 == 0 {
				if totalscolor == rl.Yellow {
					totalscolor = rl.DarkGray
				} else {
					totalscolor = rl.Yellow
				}
			}
		}
	}
}
func startlevelname() { // MARK: startlevelname

	kills = 0
	collisions = 0
	score = 0
	cloudson = true

	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Maroon, startlevelnamefade))

	if framecount%3 == 0 {
		for a := 0; a < 20; a++ {
			rl.DrawPixel(rInt(50, monw-50), rInt(50, monh-50), rl.Black)
			rl.DrawRectangle(rInt(50, monw-50), rInt(50, monh-50), 2, 2, rl.Black)
		}
	}

	lengthnamewidth := len(currentlevelname) * 30
	lengthnamewidth = (monw - lengthnamewidth) / 2

	rl.DrawText(currentlevelname, lengthnamewidth, monh/2-97, 60, rl.Black)
	rl.DrawText(currentlevelname, lengthnamewidth, monh/2-99, 60, rl.Yellow)
	rl.DrawText(currentlevelname, lengthnamewidth, monh/2-100, 60, rl.White)

	if framecount%3 == 0 {
		if startlevelnamefade > 0 {
			startlevelnamefade -= 0.04
		} else if startlevelnamefade <= 0 {
			startlevelon = false
			stoptimer = false
			pauseon = false
			startlevelnamefade = 1.0
		}
	}

}
func bosslevel() { // MARK: bosslevel
	if bossbackfadecomplete == false {
		if framecount%2 == 0 {
			bossbackfade += 0.1
			if bossbackfade >= 1.0 {
				bossbackfade = 1.0
				bossbackfadecomplete = true
			}
		}
		highlightedtileblock1 = nextblock + rInt(18, 126)
		highlightedtileblock2 = nextblock + rInt(126, 234)
		highlightedtileblock3 = nextblock + rInt(234, 342)
		highlightedtileblock4 = nextblock + rInt(342, 450)
		highlightfade1 = (rFloat32(2, 10) / 10)
		highlightfade2 = (rFloat32(2, 10) / 10)
		highlightfade3 = (rFloat32(2, 10) / 10)
		highlightfade4 = (rFloat32(2, 10) / 10)
	}

	if framecount%15 == 0 {
		highlightedtileblock1 = nextblock + rInt(18, 126)
		highlightedtileblock2 = nextblock + rInt(126, 234)
		highlightedtileblock3 = nextblock + rInt(234, 342)
		highlightedtileblock4 = nextblock + rInt(342, 450)
		highlightfade1 = (rFloat32(2, 10) / 10)
		highlightfade2 = (rFloat32(2, 10) / 10)
		highlightfade3 = (rFloat32(2, 10) / 10)
		highlightfade4 = (rFloat32(2, 10) / 10)
	}

}
func createlevel() { // MARK: createlevel

	count := 0
	drawx := 0
	drawy := 0
	blocktotal := 0

	for a := 0; a < blocknumber; a++ {

		blockxy := isoblock{}
		blockxy.xy = rl.NewVector2(float32(drawx), float32(drawy+(blockh/2)))           // left point
		blockxy.xy2 = rl.NewVector2(float32(drawx+(blockw/2)), float32(drawy))          // top point
		blockxy.xy3 = rl.NewVector2(float32(drawx+(blockw)), float32(drawy+(blockh/2))) // right point
		blockxy.xy4 = rl.NewVector2(float32(drawx+(blockw/2)), float32(drawy+(blockh))) // bottom point
		blockxy.topleft = rl.NewVector2(float32(drawx), float32(drawy))
		blockmap[blocktotal] = blockxy

		// add obstruction
		if rolldice()+rolldice()+rolldice() == 18 {
			obstructionmap[blocktotal] = true
			obstructiontypemap[blocktotal] = rInt(1, 5)
		}

		blockxy = isoblock{}
		blockxy.xy = rl.NewVector2(float32(drawx+(blockw/2)), float32(drawy+(blockh)))           // left point
		blockxy.xy2 = rl.NewVector2(float32(drawx+(blockw)), float32(drawy+(blockh/2)))          // top point
		blockxy.xy3 = rl.NewVector2(float32(drawx+(blockw+(blockw/2))), float32(drawy+(blockh))) // right point
		blockxy.xy4 = rl.NewVector2(float32(drawx+(blockw)), float32(drawy+(blockh+(blockh/2)))) // bottom
		blockxy.topleft = rl.NewVector2(float32(drawx+(blockw/2)), float32(drawy+(blockh/2)))
		blockmap[blocktotal+17] = blockxy

		// add obstruction
		if rolldice()+rolldice()+rolldice() == 18 {
			obstructionmap[blocktotal+17] = true
			obstructiontypemap[blocktotal+17] = rInt(1, 5)
		}

		blocktotal++
		count++
		drawx += 132

		if count == horizcount {
			blocktotal += 17
			count = 0
			drawx = 0
			drawy += 66
		}
	}

	leveltype()
}
func levelname(name string) { // MARK: levelname

	switch terraintype {
	case "palmisland":
		choose := rInt(0, len(tropicalnames))
		string1 := tropicalnames[choose]
		choose = rInt(0, len(woodsnames))
		string2 := woodsnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "blockforest":
		choose := rInt(0, len(rednames))
		string1 := rednames[choose]
		choose = rInt(0, len(woodsnames))
		string2 := woodsnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "brownstone":
		choose := rInt(0, len(stonenames))
		string1 := stonenames[choose]
		choose = rInt(0, len(fieldnames))
		string2 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "golf":
		choose := rInt(0, len(greennames))
		string1 := greennames[choose]
		choose = rInt(0, len(sportsnames))
		string2 := sportsnames[choose]
		choose = rInt(0, len(fieldnames))
		string3 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2 + " " + string3
	case "magenta":
		choose := rInt(0, len(magentanames))
		string1 := magentanames[choose]
		choose = rInt(0, len(fieldnames))
		string2 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "white":
		choose := rInt(0, len(whitenames))
		string1 := whitenames[choose]
		choose = rInt(0, len(fieldnames))
		string2 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "rocky":
		choose := rInt(0, len(coldnames))
		string1 := coldnames[choose]
		choose = rInt(0, len(bluenames))
		string2 := bluenames[choose]
		choose = rInt(0, len(fieldnames))
		string3 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2 + " " + string3
	case "city":
		choose := rInt(0, len(sadnames))
		string1 := sadnames[choose]
		choose = rInt(0, len(townnames))
		string2 := townnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "urban":
		choose := rInt(0, len(oddnames))
		string1 := oddnames[choose]
		choose = rInt(0, len(oasisnames))
		string2 := oasisnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "meadow":
		choose := rInt(0, len(greennames))
		string1 := greennames[choose]
		choose = rInt(0, len(fieldnames))
		string2 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "woods":
		choose := rInt(0, len(greennames))
		string1 := greennames[choose]
		choose = rInt(0, len(woodsnames))
		string2 := woodsnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "red":
		choose := rInt(0, len(rednames))
		string1 := rednames[choose]
		choose = rInt(0, len(fieldnames))
		string2 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "forest":
		choose := rInt(0, len(oddnames))
		string1 := oddnames[choose]
		choose = rInt(0, len(greennames))
		string2 := greennames[choose]
		choose = rInt(0, len(woodsnames))
		string3 := woodsnames[choose]
		currentlevelname = "the " + string1 + " " + string2 + " " + string3
	case "altforest":
		choose := rInt(0, len(oddnames))
		string1 := oddnames[choose]
		choose = rInt(0, len(psychedelicnames))
		string2 := psychedelicnames[choose]
		choose = rInt(0, len(woodsnames))
		string3 := woodsnames[choose]
		currentlevelname = "the " + string1 + " " + string2 + " " + string3
	case "mars1":
		choose := rInt(0, len(oddnames))
		string1 := oddnames[choose]
		choose = rInt(0, len(rednames))
		string2 := rednames[choose]
		choose = rInt(0, len(fieldnames))
		string3 := fieldnames[choose]
		currentlevelname = "the " + string1 + " " + string2 + " " + string3
	case "mars2":
		choose := rInt(0, len(rednames))
		string1 := rednames[choose]
		choose = rInt(0, len(basenames))
		string2 := basenames[choose]
		currentlevelname = "the " + string1 + " " + string2
	case "mars3":
		choose := rInt(0, len(oddnames))
		string1 := oddnames[choose]
		choose = rInt(0, len(rednames))
		string2 := rednames[choose]
		choose = rInt(0, len(basenames))
		string3 := basenames[choose]
		currentlevelname = "the " + string1 + " " + string2 + " " + string3
	case "mars4":
		choose := rInt(0, len(oddnames))
		string1 := oddnames[choose]
		choose = rInt(0, len(basenames))
		string2 := basenames[choose]
		currentlevelname = "the " + string1 + " " + string2
	}

}
func leveltype() { // MARK: leveltype
	choose := rInt(1, 19)
	if currentlevelnumber < 4 {
		choose = rInt(1, 8)
	}
	if currentlevelnumber > 3 && currentlevelnumber < 7 {
		choose = rInt(8, 11)
	}
	if currentlevelnumber > 6 && currentlevelnumber < 10 {
		choose = rInt(11, 15)
	}
	if currentlevelnumber > 10 && currentlevelnumber < 13 {
		choose = rInt(15, 19)
	}
	if currentlevelnumber > 13 {
		choose = rInt(1, 19)
	}

	switch choose {
	case 1:
		terraintype = "red"
		gridon = true
	case 2:
		terraintype = "white"
		gridon = true
	case 3:
		terraintype = "magenta"
		gridon = false
	case 4:
		terraintype = "woods"
		gridon = false
	case 5:
		terraintype = "urban"
		gridon = false
	case 6:
		terraintype = "meadow"
		gridon = false
	case 7:
		terraintype = "golf"
		gridon = false
	case 8:
		terraintype = "brownstone"
		gridon = false
	case 9:
		terraintype = "blockforest"
		gridon = false
	case 10:
		terraintype = "palmisland"
		gridon = false
	case 11:
		terraintype = "city"
		gridon = false
	case 12: // blue rocks
		terraintype = "rocky"
		gridon = false
	case 13:
		terraintype = "forest"
		gridon = false
	case 14:
		terraintype = "altforest"
		gridon = false
	case 15:
		terraintype = "mars1"
		gridon = false
	case 16:
		terraintype = "mars2"
		gridon = false
	case 17:
		terraintype = "mars3"
		gridon = false
	case 18:
		terraintype = "mars4"
		gridon = false
	}

	levelname(terraintype)

	switch terraintype {
	case "golf":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 9)
			if rolldice()+rolldice() < 12 {
				choose = rInt(1, 5)
			} else {
				choose = rInt(5, 9)
			}
			blocktiles[a] = choose
		}
	case "magenta":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 9)
			blocktiles[a] = choose
		}
	case "white":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 9)
			blocktiles[a] = choose
		}
	case "blockforest":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 10)
			if rolldice()+rolldice() < 10 {
				choose = rInt(1, 5)
			} else {
				choose = rInt(5, 10)
			}
			blocktiles[a] = choose
		}
	case "palmisland":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 11)
			if rolldice()+rolldice() < 11 {
				choose = rInt(1, 5)
			} else {
				choose = rInt(5, 11)
			}
			blocktiles[a] = choose
		}
	case "brownstone":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 8)
			blocktiles[a] = choose
		}
	case "mars4":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 39)
			blocktiles[a] = choose
		}
	case "mars3":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 39)
			blocktiles[a] = choose
		}
	case "mars2":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 39)
			blocktiles[a] = choose
		}
	case "mars1":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 21)
			blocktiles[a] = choose
		}
	case "altforest":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 21)
			blocktiles[a] = choose
		}
		randomcolors()

		altforestcolor1 = randomcolorslice[0]
		altforestcolor2 = randomcolorslice[1]
		altforestcolor3 = randomcolorslice[2]
		altforestcolor4 = randomcolorslice[3]
		altforestcolor5 = randomcolorslice[4]
		altforestcolor6 = randomcolorslice[5]
		altforestcolor7 = randomcolorslice[6]
		altforestcolor8 = randomcolorslice[7]
		altforestcolor9 = randomcolorslice[8]
		altforestcolor10 = randomcolorslice[9]

	case "forest":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 21)
			blocktiles[a] = choose
		}
	case "rocky":
		for a := 0; a < blocknumber; a++ {
			blocktiles[a] = 1
		}
		for a := 0; a < blocknumber; a++ {
			if rolldice()+rolldice() > 9 {
				blocktiles[a] = rInt(2, 9)
			}
		}
	case "city":
		for a := 0; a < blocknumber; a++ {
			blocktiles[a] = 1
		}
		for a := 0; a < blocknumber; a++ {
			if rolldice()+rolldice() > 9 {
				blocktiles[a] = rInt(2, 13)
			}
		}
	case "urban":
		for a := 0; a < blocknumber; a++ {
			blocktiles[a] = 1
		}
		for a := 0; a < blocknumber; a++ {
			if rolldice() > 5 {
				blocktiles[a] = rInt(2, 5)
			}
		}
	case "meadow":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 9)
			blocktiles[a] = choose
		}
	case "woods":
		for a := 0; a < blocknumber; a++ {
			blocktiles[a] = 1
		}
		for a := 0; a < blocknumber; a++ {
			if rolldice() > 4 {
				blocktiles[a] = rInt(2, 12)
			}
		}
	case "red":
		for a := 0; a < blocknumber; a++ {
			choose := rInt(1, 9)
			blocktiles[a] = choose
		}
	}

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "isometric")
	setscreen()
	rl.CloseWindow()
	initialize()
	createlevel()
	raylib()
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw-300, 0, 500, monw, rl.Fade(rl.Blue, 0.4))
	rl.DrawFPS(monw-290, monh-100)

	mousexTEXT := fmt.Sprintf("%g", mousepos.X)
	mouseyTEXT := fmt.Sprintf("%g", mousepos.Y)
	vertcountTEXT := strconv.Itoa(vertcount)
	horizcountTEXT := strconv.Itoa(horizcount)
	blocknumberTEXT := strconv.Itoa(blocknumber)
	mouseblockTEXT := strconv.Itoa(mouseblock)
	checkblock := blockmap[nextblock]
	checkblockxTEXT := fmt.Sprintf("%g", checkblock.topleft.X)
	checkblockyTEXT := fmt.Sprintf("%g", checkblock.topleft.Y)
	cameracloudszoomTEXT := fmt.Sprintf("%g", cameraclouds.Zoom)
	cameracloudsyTEXT := fmt.Sprintf("%g", cameraclouds.Target.Y)
	cameraxTEXT := fmt.Sprintf("%g", camera.Target.X)
	camerayTEXT := fmt.Sprintf("%g", camera.Target.Y)
	playerxTEXT := fmt.Sprintf("%g", playerxy.X)
	playeryTEXT := fmt.Sprintf("%g", playerxy.Y)
	playerhpTEXT := strconv.Itoa(playerhp)
	highlightfade1TEXT := fmt.Sprintf("%g", highlightfade1)
	bulletmodifierTEXT := fmt.Sprintf("%g", bulletmodifiers.number/2)
	bulletxycheckxTEXT := fmt.Sprintf("%g", bulletxycheck.X)
	framecountTEXT := strconv.Itoa(framecount)
	checkcoin2 := coinsmap[1]
	coinTEXT := fmt.Sprintf("%g", checkcoin2.rec.Y)
	currentlevelnumberTEXT := strconv.Itoa(currentlevelnumber)
	greenshippurchasedTEXT := strconv.FormatBool(greenshippurchased)
	storeselectTEXT := strconv.Itoa(storeselect)
	nextlevelselectTEXT := strconv.Itoa(nextlevelselect)
	cloudsonTEXT := strconv.FormatBool(cloudson)

	rl.DrawText(mousexTEXT, monw-290, 10, 10, rl.White)
	rl.DrawText("mouseX", monw-150, 10, 10, rl.White)
	rl.DrawText(mouseyTEXT, monw-290, 20, 10, rl.White)
	rl.DrawText("mouseY", monw-150, 20, 10, rl.White)
	rl.DrawText(vertcountTEXT, monw-290, 30, 10, rl.White)
	rl.DrawText("vertcount", monw-150, 30, 10, rl.White)
	rl.DrawText(horizcountTEXT, monw-290, 40, 10, rl.White)
	rl.DrawText("horizcount", monw-150, 40, 10, rl.White)
	rl.DrawText(blocknumberTEXT, monw-290, 50, 10, rl.White)
	rl.DrawText("blocknumber", monw-150, 50, 10, rl.White)
	rl.DrawText(checkblockxTEXT, monw-290, 60, 10, rl.White)
	rl.DrawText("checkblockx", monw-150, 60, 10, rl.White)
	rl.DrawText(checkblockyTEXT, monw-290, 70, 10, rl.White)
	rl.DrawText("checkblocky", monw-150, 70, 10, rl.White)
	rl.DrawText(mouseblockTEXT, monw-290, 80, 10, rl.White)
	rl.DrawText("mouseblock", monw-150, 80, 10, rl.White)
	rl.DrawText(cameracloudszoomTEXT, monw-290, 90, 10, rl.White)
	rl.DrawText("cameracloudszoom", monw-150, 90, 10, rl.White)
	rl.DrawText(cameracloudsyTEXT, monw-290, 100, 10, rl.White)
	rl.DrawText("cameracloudsy", monw-150, 100, 10, rl.White)
	rl.DrawText(cameraxTEXT, monw-290, 110, 10, rl.White)
	rl.DrawText("camerax", monw-150, 110, 10, rl.White)
	rl.DrawText(camerayTEXT, monw-290, 120, 10, rl.White)
	rl.DrawText("cameray", monw-150, 120, 10, rl.White)
	rl.DrawText(playerxTEXT, monw-290, 130, 10, rl.White)
	rl.DrawText("playerx", monw-150, 130, 10, rl.White)
	rl.DrawText(playeryTEXT, monw-290, 140, 10, rl.White)
	rl.DrawText("playery", monw-150, 140, 10, rl.White)
	rl.DrawText(playerhpTEXT, monw-290, 150, 10, rl.White)
	rl.DrawText("playerhp", monw-150, 150, 10, rl.White)
	rl.DrawText(highlightfade1TEXT, monw-290, 160, 10, rl.White)
	rl.DrawText("highlightfade1", monw-150, 160, 10, rl.White)
	rl.DrawText(bulletmodifierTEXT, monw-290, 170, 10, rl.White)
	rl.DrawText("bulletmodifier", monw-150, 170, 10, rl.White)
	rl.DrawText(bulletxycheckxTEXT, monw-290, 180, 10, rl.White)
	rl.DrawText("bulletxycheckx", monw-150, 180, 10, rl.White)
	rl.DrawText(framecountTEXT, monw-290, 190, 10, rl.White)
	rl.DrawText("framecount", monw-150, 190, 10, rl.White)
	rl.DrawText(coinTEXT, monw-290, 200, 10, rl.White)
	rl.DrawText("checkcoin2recY", monw-150, 200, 10, rl.White)
	rl.DrawText(currentlevelnumberTEXT, monw-290, 210, 10, rl.White)
	rl.DrawText("currentlevelnumber", monw-150, 210, 10, rl.White)
	rl.DrawText(greenshippurchasedTEXT, monw-290, 220, 10, rl.White)
	rl.DrawText("greenshippurchased", monw-150, 220, 10, rl.White)
	rl.DrawText(storeselectTEXT, monw-290, 230, 10, rl.White)
	rl.DrawText("storeselect", monw-150, 230, 10, rl.White)
	rl.DrawText(nextlevelselectTEXT, monw-290, 240, 10, rl.White)
	rl.DrawText("nextlevelselect", monw-150, 240, 10, rl.White)
	rl.DrawText(cloudsonTEXT, monw-290, 250, 10, rl.White)
	rl.DrawText("cloudson", monw-150, 250, 10, rl.White)

}
func clearmaps() { // MARK: clearmaps

	for a := 0; a < blocknumber*2; a++ {
		obstructiontypemap[a] = 0
		obstructionmap[a] = false
		blocktiles[a] = 0
		blockmap[a] = isoblock{}
	}
	for a := 0; a < len(bulletmap); a++ {
		bulletmap[a] = bullet{}
	}
	for a := 0; a < len(powerupsmap); a++ {
		powerupsmap[a] = powerup{}
	}
	for a := 0; a < len(coinsmap); a++ {
		coinsmap[a] = coinstruct{}
	}
}
func createmaps() { // MARK: createmaps
	obstructiontypemap = make([]int, blocknumber*2)
	obstructionmap = make([]bool, blocknumber*2)
	blocktiles = make([]int, blocknumber*2)
	blockmap = make([]isoblock, blocknumber*2)
}
func initialize() { // MARK: initialize

	vertcount = (monh / 66)
	horizcount = (monw / 132) + 3
	gridlayout = horizcount * vertcount
	screenblocknumber = horizcount*(vertcount*2) + (horizcount * 4)
	blocknumber = horizcount * ((vertcount * 2) * 100)
	nextblock = blocknumber - screenblocknumber*2

	createmaps()
	playerx = float32(monw / 2)
	playery = float32(monh / 2)
	playerxy = rl.NewVector2(playerx, playery)

	scanlineson = true
	cloudson = true

	startscreenship = rl.NewRectangle(float32(monw/2), float32(monh/2), 112, 75)
	poweruptimers()
	createcoins()
	currentlevelnumber = 1
	playership = ship1
	scannercolor = rl.DarkBlue
	currenttarget = lasermarker1
	pixelnoiseon = true
	bulletmodifiers.img = 1
	createweather()

}
func setscreen() { // MARK: setscreen
	monh = rl.GetScreenHeight()
	monw = rl.GetScreenWidth()
	monh32 = int32(monh)
	monw32 = int32(monw)
	rl.SetWindowSize(monw, monh)
	camera.Zoom = 1.0
	camera.Target.X = 66
	camera.Target.Y = 33

	cameraclouds.Zoom = 2.0
	camera4X.Zoom = 4.0
	camera2X.Zoom = 2.0
	camera15X.Zoom = 1.5

} // random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
