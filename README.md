begin

    Obj Player
        String name
        Piece pieces[]
    
    Obj Piece
        String colour
        String name
        String type
        Point position
        Point moves[]

    Obj Board
        Point squares[] 

    Obj Point
        String letter
        String number
        Piece occupied


    setupGraphics()

    Player White 
    Player Black
    Point Board[] = creatBoard()
    Player currentPlayer = White

    Setup()


    while true
        Piece p = randPickPiece(currentPlayer)
        randMove(legalMoves(p), board)
        graphicUpdate(board)
        if winCondition 
            print winGame
            break
        else if stalemate 
            print drawGame
            break
        end if
        swapCurrentPlayer
    end while
end



begin func setup
    White = newPlayer("white")
    Black = newPlayer("black")
end

begin func legalMoves(Piece p)
    Point returnMoves[]
    if p.type not "knight"
        for dir, dist in direction(p) # //TO-DO direction implementation
            int step = 1
            Point nextStep
            while step <= dist
                switch dir
                    case "nw"
                        nextStep = Point(p.position.letter - step, p.position.number + step))
                    case "n"
                        nextStep = Point(p.position.letter, p.position.number + step))
                    case "ne"
                        nextStep = Point(p.position.letter + step, p.position.number + step))
                    case "w"
                        nextStep = Point(p.position.letter - step, p.position.number))
                    case "e"
                        nextStep = Point(p.position.letter + step, p.position.number))
                    case "sw"
                        nextStep = Point(p.position.letter - step, p.position.number - step))
                    case "s"
                        nextStep = Point(p.position.letter, p.position.number - step))
                    case "se"
                        nextStep = Point(p.position.letter + step, p.position.number - step))
                end switch
                switch check(nextStep) //TO-DO implement check(Point)
                    case "empty"
                        returnMoves.add(nextStep)
                    case "capture"
                        returnMoves.add(nextStep)
                        end while
                    case "friend"
                        end while
                end switch     
                step+=1
            end while
        end for
    else 
        //knight moves
    end if
    return returnMoves[]
end

begin func randPickPiece(Player p)
    return rand(p.pieces)
end

being func swapCurrentPlayer
        sleep
end
        
begin func randMove

end

begin legalStep(Point newPoint)

end



