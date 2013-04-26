###################################################
## Algorithm for the deferred acceptance problem ##
###################################################

import heapq


## Objects representing the agents involved: acceptor and proposor. ##

class Acceptor:
  "acceptor with ranked list of proposors"

  def __init__(self, name):
    self.proposors = [] #heapq
    self.name = name
    self.accepted_proposor = None

  def get_proposors(self):
    return self.proposors

  def set_proposors(self, ranked_proposors):
    self.proposors = ranked_proposors

  def get_accepted_proposor(self):
    return self.accepted_proposor

  def set_accepted_proposor(self, proposor):
    if (self.accepted_proposor): #replace previously accepted proposor
      replaced = self.accepted_proposor
      self.accepted_proposor = proposor
      return true, replaced
    else:
      self.accepted_proposor = proposor
      return false, None #otherwise, nothing to replace

  def get_name(self):
    return self.name


class Proposor:
  "acceptor with ranked list of acceptors"

  def __init__(self, name):
    self.acceptors = [] #heapq
    self.name = name

  def get_acceptors(self):
    return self.acceptors
  
  def get_highest_acceptor(self):
    return heapq.heappop(acceptors)

  def set_acceptors(self, ranked_acceptors):
    self.acceptors = ranked_acceptors

  def get_name(self):
    return self.name


### Functions for solving deferred acceptance problem with n acceptors, n proposors ###

def propose(proposor, unmatched_proposors):
  highest_acceptor = proposor.get_highest_acceptor()
  competition = highest_acceptor.get_accepted_proposor()
  if not competiton or (competition and competition > propsor):
    didReplace, replaced = highest_acceptor.set_accepted_proposor(proposor)
    if didReplace:
      unmatched_proposors.append(replaced) #add replaced to storage for proposal pool for next round
  else:
    unmatched_proposors.append(proposor) #add back to storage for proposal pool for next round

def deferred_acceptance_algorithm(acceptors, proposors):
  proposal_pool = [proposor for proposor in proposors]
  current_round = 0
  while len(proposal_pool) > 0:
    current_round += 1
    unmatched_proposors = []
    for proposor in proposal_pool:
      propose(proposor, unmatched_proposors)
    proposal_pool = unmatched_proposors


##################################################################
# Initializing a simple example for 4 acceptors and 4 proposors. #
##################################################################

def initialize_acceptor(acceptor, ranks, proposors):
  heap = []
  for i in range(0, len(ranks)):
    heapq.heappush(heap, (ranks[i], proposors[i]))
  acceptor.set_proposors(heap)

def initialize_proposor(proposor, ranks, acceptors):
  heap = []
  for i in range(0, len(ranks)):
    heapq.heappush(heap, (ranks[i], acceptors[i]))
  proposor.set_acceptors(heap)


acceptor_1 = Acceptor("A")
acceptor_2 = Acceptor("B")
acceptor_3 = Acceptor("C")
acceptor_4 = Acceptor("D")
acceptors = [ acceptor_1, acceptor_2, acceptor_3, acceptor_4 ]

proposor_1 = Proposor("1")
proposor_2 = Proposor("2")
proposor_3 = Proposor("3")
proposor_4 = Proposor("4")
proposors = [ proposor_1, proposor_2, proposor_3, proposor_4 ]

proposor_1_ranks = [ 2, 1, 3, 4 ]
proposor_2_ranks = [ 2, 3, 4, 1 ]
proposor_3_ranks = [ 1, 3, 2, 4 ]
proposor_4_ranks = [ 2, 3, 1, 4 ]

acceptor_1_ranks = [ 1, 3, 2, 4 ]
acceptor_2_ranks = [ 3, 4, 1, 2 ]
acceptor_3_ranks = [ 4, 2, 3, 1 ]
acceptor_4_ranks = [ 3, 2, 1, 4 ]

initialize_acceptor(acceptor_1, acceptor_1_ranks, proposors)
initialize_acceptor(acceptor_2, acceptor_2_ranks, proposors)
initialize_acceptor(acceptor_3, acceptor_3_ranks, proposors)
initialize_acceptor(acceptor_4, acceptor_4_ranks, proposors)

initialize_proposor(proposor_1, proposor_1_ranks, acceptors)
initialize_proposor(proposor_2, proposor_2_ranks, acceptors)
initialize_proposor(proposor_3, proposor_3_ranks, acceptors)
initialize_proposor(proposor_4, proposor_4_ranks, acceptors)

deferred_acceptance_algorithm(acceptors, proposors)
for acceptor in acceptors:
  print "A: " + acceptor.get_name + "-> P: " + acceptor.get_accepted_proposor()
