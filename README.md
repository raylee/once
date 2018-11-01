Once is what I usually want from uniq -- preserve order but don't show
duplicates.

Usage:

    once [file ...]

Once reads the named input files a line at a time. If the line isn't
in a map[string]bool, it's printed and added to the map.
