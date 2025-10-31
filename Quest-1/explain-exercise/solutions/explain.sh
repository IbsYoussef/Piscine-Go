#!/bin/sh

# Commands: echo
# Documents the investigation process by outputting key evidence
#
# Investigation Summary:
#
# Line 1: Key witness name
#   From crimescene file: Barista mentioned "Annabel" left before shots
#   grep "Annabel" mystery/people → Found multiple Annabel entries
#   Followed address to find the correct Annabel Church
#
# Line 2: Interview number
#   sed -n '179p' mystery/streets/Buckingham_Place
#   Result: "SEE INTERVIEW #699607"
#   This interview contains the car description
#
# Line 3: Suspect's car description
#   grep -r "699607" mystery/interviews/
#   Interview states: BLUE HONDA with license plate L337...9
#
# Lines 4-6: Other main suspects (alphabetically by last name)
#   Step 1: Find all people in all four memberships
#     comm -12 <(sort mystery/memberships/AAA) <(sort mystery/memberships/Delta_SkyMiles) | \
#     comm -12 - <(sort mystery/memberships/Terminal_City_Library) | \
#     comm -12 - <(sort mystery/memberships/Museum_of_Bash_History)
#
#   Step 2: Filter for tall males (≥ 6 feet)
#     grep -Ff suspects.txt mystery/people | awk -F '\t' '$2=="M" && $3>=72'
#
#   Step 3: Cross-reference with vehicle owner
#     awk -v RS='' '/License Plate L337.*9/ && /Make: Honda/ && /Color: Blue/' mystery/vehicles
#     Result: Joe Germuska (arrested)
#
#   Step 4: List remaining suspects, sorted by last name
#     Results: Aldo Nicolas, Andrei Masna, Matt Waite
#
# Output format:
# Witness full name
# Interview number
# Car color and make
# Suspect 1 (by last name)
# Suspect 2 (by last name)
# Suspect 3 (by last name)

echo "Annabel Church"
echo "699607"
echo "Blue Honda"
echo "Aldo Nicolas"
echo "Andrei Masna"
echo "Matt Waite"