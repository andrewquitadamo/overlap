#usr/bin/perl
use strict;
use warnings;

my @files = @ARGV;

my @index;
my $header;
my @header;
my %id_hash;
my @id_array;

foreach (@files)
{
	open(FILE,$_) || die "Can't open file $_.\n";
	my $header=<FILE>;
	chomp $header;
	my @header = split('\t',$header);
	my $i=0;
	foreach(@header)
	{
		if (exists $id_hash{$_})
		{
			$id_hash{$_}=$id_hash{$_}.",".$i;
		}
		else
		{
			$id_hash{$_}=$i;
		}
		$i++;
	}
	close FILE;
}

foreach (@files)
{
	my @pos;
	
	open(FILE,$_) || die "Can't open file $_.\n";
	close FILE;
}

#my @keys = keys %id_hash;
#foreach(@keys)
#{
#	print "$_\t$id_hash{$_}\n";
#}


#Read sample IDs into an array
#open(FILE,$sample_file) || die "Can't open file $sample_file.\n";
#my @samples=<FILE>;
#close FILE;

#open FILE,">"."$vcf_file" . ".samples" or die $!;
#print FILE $output;
#close FILE;
