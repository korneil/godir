#/bin/bash

function godir {
  IFS=$'\n'

  o=$(./godir $@);

  if [ -z "$o" ]; then
    echo "No match."
  else
    nl=`echo "$o" | wc -l`;
    if [ $nl -eq 1 ]; then
      echo "$o"
    else
      id=0;
      for dir in $o; do
        echo $id: $dir;
        id=`expr $id + 1`
      done
      read -p "Select: " s;
      a=($o)
      echo "${a[s]}";
    fi
  fi
}

godir ~/Remote $@