with import <nixpkgs> {};

stdenv.mkDerivation rec {
	name = "terraform-provider-auth0";
	env = buildEnv { name = name; paths = buildInputs; };

	buildInputs =
		[
			gnumake
			terraform_0_11
			go
		];
}
