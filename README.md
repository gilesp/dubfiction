dubfiction
==========

A tool to aid in dub fiction remixes, a la Jeff Noon: https://web.archive.org/web/20160304064129/http://www.metamorphiction.com/index.php/the-ghost-on-the-b-side-remixing-narrative/

Written (badly) in go. Install it with a go get github.com/gilesp/dubfiction or similar command

usage
=====

At it's simplest, just run "dubfiction" and it will look for a text file called main.txt and perform a first run dub remix of it. It uses a random phrase size of between 2 and 7 words.

You can specify a different file to use with the -main option:

    dubfiction -main=main_file_to_remix

If you want to mix in a secondary text, then call it like so:

    dubfiction -main=main_file_to_remix -secondary=file_to_mix_in

the secondary text will be mixed in randomly.

If you want to limit the output to approximately 160chars, then add the -tweet=true option.


Take the output, modify it as per your desires then run it back through dubfiction for a second pass.

advanced
========

Running on linux|os x|similar? don't feel like manually modding the output and want to do two passes in one?

    dubfiction -main=blah.txt -secondary=flavouring.txt > phase_one.txt && dubfiction -main=phase_one.txt

savour the aftertaste.
