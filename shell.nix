{ pkgs ? import<nixpkgs> {}}:

with pkgs;

mkShell {
	SHELL_NAME = "SemanticVectors";
	buildInputs = [ go ];
	shellHook = ''
	export PATH=$HOME/go/bin:$PATH
	'';
}
