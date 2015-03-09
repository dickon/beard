# beard; an experimental off-line deduplication system


## Motivation

I have a Linux file server which stores backups in various formats for
various machines I have. It is quite full and has spare compute
time. I'm interested to see how many more backups I can store in traditional
formats if I use a filesystem such as btrfs which can share extents between files.

### RAM consumption

duperemove, ZFS and the btrfs patches for online deduplication in the kernel
store their hash tables in RAM. That's probably a limiting factor given
the size of disks. Databases servers should be good enough to make good use of
RAM to cache data access.

### Offline versus inline

Inline dedup may have it uses but sometimes simple data paths work well. There
are enough facilities in modern operating systems for detecting changes that we
should be able to run off line and avoid walking the whole filesystem. So at least
to start with we should allow writes to happen as normal and then spend idle compute
time looking for duplicated information. 

### Collisions

We'll end up wanting to consider every possible byte offset within each
file for deduplication, to cope with prepends. I don't know of an efficient way
to do that with a secure hash function. I also care about the hash value size,
so let's not assume that our first pass hash function has a neglible risk of collisions;
instead do the rolling adler32, optionally then a user configurable secure hash and verify
our assumptions by doing an equality test on the identified duplciate data.


## Design

The main datastructure is a postgresql database with these tables:

adlerhashes table with columns:
  checksum INTEGER NOT NULL,
  blocknum INTEGER FOREIGN KEY REFERENCS blocks(blocknum)
 
blocks table with columns:
  blocknum INTEGER NOT NULL,
  locationnum  INTEGER NOT NULL FOREIGN KEY REFERENCES locations(locationnum)

locations table
  locationnum INTEGER NOT NULL
  pathnum INTEGER NOT NOLL FOREIGN KEY REFERENCES paths(pathnum)
  offset BIGINT
  length INTEGER
  securehash BYTEA
 
paths table
  pathnum INTEGER PRIMARY KEY NOT NULL
  mtime INTEGER -- mtime of file when we read it
  readtime INTEGER -- time we read the file
  inode INTEGER
  path TEXT -- path relative to filesystem top
  filesystemnum INTEGER NOT NULL FOREIGN KEY

filesystems table
  filesystemnum INTEGER NOT NULL PRIMARY KEY
  mount TEXT


### implementation language

+ some operations are going to take ages. Getting a python NameError due to a typo at the end 
  would be annoying, so some compile time checking is desirable.
+ there's likely to be some bit twiddling
+ both the interactions with the database and the filesystem will be slow 
  enough that being able to pipeline them is useful.