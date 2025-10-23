#!/bin/sh
# Exercise: explain
#
# Optional:
# The commissioner wants to know how you solved the mystery.
# This script prints, in order:
# 1. The first and last name of your key witness
# 2. The interview number of this witness
# 3. The colour and make of the car of the main suspect
# 4. The names of the three other main suspects not arrested
#    (sorted alphabetically by their last name)
#
# Explanation:
# - Witness identified from crime scene clues → Annabel Church
# - Related interview number → 699607 (from Buckingham Place)
# - Suspect’s car → Blue Honda
# - Other main suspects derived from the narrowed list of tall males
#   sharing all memberships, excluding the arrested Joe Germuska.
#
# Example output:
# $ ./explain.sh | cat -e
# Annabel Church$
# 699607$
# Blue Honda$
# <suspect1>$
# <suspect2>$
# <suspect3>$
# $
#
# Solution:
echo "Annabel Church"
echo "699607"
echo "Blue Honda"
echo "Aldo Nicolas"
echo "Andrei Masna"
echo "Matt Waite"