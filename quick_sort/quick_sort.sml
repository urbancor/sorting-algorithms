(* Seznam uredi naraščajoče z algoritmom quicksort. Prvi argument je funkcija za primerjanje. *)
(* Sorts the list with quicksort. First argument is a compare function. *)
(*fn : ('a * 'a -> order) -> 'a list -> 'a list*)
fun quicksort f s =
    case s of
    [] => []
    | h::t => 
        let
            val left = List.filter (fn x => f(x,h) = LESS) t;
            val right = List.filter (fn x => f(x,h) = GREATER) t;
            val middle = List.filter (fn x => f(x,h) = EQUAL) t;
            val middle = h::middle;
        in
            quicksort f left @ middle @ quicksort f right
        end
