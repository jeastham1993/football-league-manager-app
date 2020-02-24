namespace LeagueManager.Models
{
    public class TeamSearchResponse
    {
        public Team[] Teams { get; set; }

        public Team Team { get; set; }

        public string Err { get; set; }
    }
}