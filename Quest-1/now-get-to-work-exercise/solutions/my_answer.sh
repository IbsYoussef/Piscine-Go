#!/bin/sh
# Exercise: now-get-to-work
#
# Optional:
# Clone and investigate the repo:
#   https://github.com/01-edu/the-final-cl-test
# Your task is to work from the provided files (no internet) to identify the
# murderer and print ONLY their full name (plus a newline).
#
# The final-cl-test problem statement (summary):
# ------------------------------------
# "Something terrible happened." You must solve a command-line murder mystery.
# Use the files under `mystery/` to gather clues:
# - `crimescene`: narrative with marked CLUEs
# - `memberships/`: membership rosters (plain names)
# - `people`: tab-delimited demographics (Name, Gender, Age, Address…)
# - `streets/`: one file per street; each line is a house entry or a pointer
# - `interviews/`: text transcripts, sometimes referenced by an ID
# - `vehicles`: one vehicle per paragraph (blank line separated)
#
# Repo-specific clue breakdown:
# -----------------------------
# From `crimescene` (grep for CLUE):
#   • Perp is a tall male (≥ 6').
#   • Wallet had AAA, Delta SkyMiles, Terminal City Library,
#     Museum of Bash History membership cards.
#   • Barista: a woman named Annabel left just before the shots.
#
# Following the Annabel lead:
#   • In `people`, there are several "Annabel …" entries with street & line.
#   • `streets/Buckingham_Place` at line 179 says: SEE INTERVIEW #699607.
#   • That interview states the fleeing car was a BLUE HONDA, license plate
#     starts with "L337" and ends with "9".
#
# Reproducible investigation steps (commands used):
# -------------------------------------------------
# 1) Collect CLUEs:
#      grep -n 'CLUE' crimescene
#
# 2) Wallet intersection (names in ALL FOUR lists):
#      comm -12 <(comm -12 <(sort memberships/AAA) <(sort memberships/Delta_SkyMiles) \
#                 | comm -12 - <(sort memberships/Terminal_City_Library)) \
#               <(sort memberships/Museum_of_Bash_History) \
#        > suspects_by_memberships.txt
#
# 3) Cross with `people` (tab-delimited), keep only males:
#      grep -Ff suspects_by_memberships.txt people | awk -F '\t' '$2=="M"' > male_suspects.txt
#
# 4) Follow the Annabel → street → interview pointer:
#      grep -RIn "Annabel" .
#      sed -n '179p' streets/Buckingham_Place     # -> SEE INTERVIEW #699607
#      grep -RIn "699607" .                       # read interview: BLUE HONDA, L337…9
#
# 5) Find the BLUE HONDA with plate L337…9 and read its owner:
#      awk -v RS='' '/License Plate L337.*9/ && /Make: Honda/ && /Color: Blue/' vehicles
#    → Owner identified as **Joe Germuska** (Height 6'2"), matching "tall male".
#
# Notes:
# - Output must be exactly one line: the culprit's full name + newline.
# - Keep tooling to standard CLI utilities available in the repo context.
#
# Example output:
# $ bash my_answer.sh | cat -e
# Joe Germuska$
# $
#
# Solution:
echo "Joe Germuska"