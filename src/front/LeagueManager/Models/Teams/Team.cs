using System.Collections.Generic;

namespace LeagueManager.Models
{
    public class Team
    {
        public string Id { get; set; }

        public string Name { get; set; }

        public IEnumerable<Player> Players { get; set; }
    }
}