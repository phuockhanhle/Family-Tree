# Family-Tree

## Terminologies:
- _Upside_ (from father to grandfather) --> rank - 1
- _Downside_ (from grandfather to father) --> rank + 1

## Desired UI
### Visualization
- **Overall**: Only people sharing DNA (biologically related) are visualized. Hidden people should be represented with specific symbols/signs.
- **Vertically**: In a view of each person (selected by clicked or first registered), only people with rank within range $\pm$ 3.
- **Horizontally**: People with same rank are only visualized if they have the same root at rank = *their_rank* - 2

## Desired UX
Cach goi
=== cung rank ==
- neu 2 nguoi same_root:
    + xet elder cua 2 nguoi o (rank(root_chung) + 1)
- neu 2 nguoi ko same_root:
    + check vo (hoac chong) co same_root ko
        - neu ko --> so sanh tuoi
        - neu co --> so elder theo vo (hoac chong)
=== rank + 1 ===
- nguoi duoi goi nguoi tren -> xet quan he cua nguoi tren voi ba/me cua nguoi duoi:
    + neu same root -> return elder + reference toi rules
    + neu ko same root -> xet quan he cua vo/chong cua nguoi tren voi ba/me cua nguoi duoi -> return elder + reference toi rules
- nguoi tren goi nguoi duoi la con
=== rank + 2 ===
- nguoi duoi goi nguoi tren la ong/ba noi/ngoai (same root), ong/ba (khac root)
- nguoi tren goi nguoi duoi chau
=== rank + 3 ===
- nguoi duoi goi nguoi tren la ong/ba co (same root), ong/ba (khac root)
- nguoi tren goi nguoi duoi la chau

## Tools:
- Language: Go, JS
- Framework: d3js/c3js, css-d3-mitch