
import Data.List
import System.IO

catcost = 10.00
fishcost = 3.00
birdcost = 0.50

totalcost = 100.00
totalnum = 100

maxcats :: Int
maxcats = truncate ( totalcost / catcost )
cats = [ ( x, fromIntegral x * catcost ) | x <- [ 1 .. maxcats] ]

maxfish :: Int
maxfish = truncate ( totalcost / fishcost )
fish = [ ( x, fromIntegral x * fishcost ) | x <- [ 1 .. maxfish ] ]

maxbirds :: Int
maxbirds = truncate ( totalcost / birdcost )
birds = [ ( x, fromIntegral x * birdcost ) | x <- [ 1 .. maxbirds ] ]

solutions = [ ( numcats, numfish, numbirds ) 
          | (numcats, costcats) <- cats,
            (numfish, costfish) <- fish,
            (numbirds, costbirds) <- birds,
            numcats + numfish + numbirds == 100,
            costcats + costfish + costbirds == 100.00 ]

main = putStrLn (show solutions)

