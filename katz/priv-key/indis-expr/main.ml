type key =
  | Kall
  | Kone
  | Ktwo

type msg =
  | Mall
  | Mone
  | Mtwo

type cpr =
  | Call
  | Cone
  | Ctwo

let rec p (k : key) (m : msg) (c : cpr) : float =
  match (k, m, c) with
  | (Kall, Mall, Call) -> 1.
  | (Kone, _, _) -> 0.5 *. (p Kall m c)
  | (_, Mone, _) -> 0.5 *. (p k Mall c)
  | (_, _, Cone) -> 0.5 *. (p k m Call)
  | _ -> 0.

let () =
  Printf.printf "%f\n" (p Kone Mone Cone)
