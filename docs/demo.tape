# Terminal showcase for the README.
# Requires network access to PokéAPI (or VPN if blocked).
# Render from the repository root with:
#   vhs docs/demo.tape

Output docs/demo.gif

Set Shell "bash"
Set FontFamily "GoogleSansCode Nerd Font Mono"
Set FontSize 21
Set LineHeight 1.3
Set LetterSpacing 0.4
Set Width 1180
Set Height 640
Set Padding 16
Set Theme "Catppuccin Mocha"
Set WindowBar "Colorful"
Set WindowBarSize 40
Set BorderRadius 12
Set TypingSpeed 45ms
Set Framerate 30
Set PlaybackSpeed 1.0

Hide
Type "clear"
Enter
Show
Sleep 500ms

Type "./pokedexcli"
Enter
Sleep 1200ms

Type "explore canalave-city-area"
Enter
Sleep 4s

Type "catch magikarp"
Enter
Sleep 7s

Type "inspect magikarp"
Enter
Sleep 3s

Type "pokedex"
Enter
Sleep 2s

Type "exit"
Enter
Sleep 2s
