using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;
using LeagueManager.Models;
using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Logging;

namespace LeagueManager.Services
{
    public class TeamState
    {
        private string LastSearch = "";

        public IReadOnlyList<Team> TeamSearchResults { get; private set; }
        public bool SearchInProgress { get; private set; }
        public event Action OnChange;

        private readonly HttpClient _httpClient;
        private readonly ILogger<TeamState> _logger;

        public TeamState(
            HttpClient httpClient,
            ILogger<TeamState> logger)
        {
            this._httpClient = httpClient;
            this._logger = logger;
        }

        public async Task DeletePlayer(Player player)
        {
            try
            {
                this._logger.LogInformation($"{player.TeamId} - {player.Name} - {player.Position}");

                var httpContent = new StringContent(JsonSerializer.Serialize(player), Encoding.UTF8, "application/json");

                var postResult = await this._httpClient.DeleteAsync($"http://localhost:8080/team/{player.TeamId}/players");

                this._logger.LogInformation(postResult.StatusCode.ToString());

                if (postResult.IsSuccessStatusCode)
                {
                    NotifyStateChanged();
                }
                else
                {
                    throw new Exception(await postResult.Content.ReadAsStringAsync());
                }
            }
            catch (Exception ex)
            {
                this._logger.LogError(ex.Message);
                this._logger.LogError("Failure adding player");
            }
        }
        public async Task AddPlayer(Player player)
        {
            try
            {
                this._logger.LogInformation($"{player.TeamId} - {player.Name} - {player.Position}");

                var httpContent = new StringContent(JsonSerializer.Serialize(player), Encoding.UTF8, "application/json");

                var postResult = await this._httpClient.PostAsync($"http://localhost:8080/team/{player.TeamId}/players", httpContent);

                this._logger.LogInformation(postResult.StatusCode.ToString());

                if (postResult.IsSuccessStatusCode)
                {
                    NotifyStateChanged();
                }
                else
                {
                    throw new Exception(await postResult.Content.ReadAsStringAsync());
                }
            }
            catch (Exception ex)
            {
                this._logger.LogError(ex.Message);
                this._logger.LogError("Failure adding player");
            }
        }
        public async Task<Team> GetSpecific(string teamId)
        {
            this._logger.LogInformation("Running HTTP search");

            var team = await this._httpClient.GetJsonAsync<TeamSearchResponse>($"http://localhost:8080/team/{teamId}");

            return team.Team;
        }
        public async Task ReRunSearch()
        {
            await this.Search(LastSearch);
        }
        public async Task Search(string searchTerm)
        {
            try
            {
                this.LastSearch = searchTerm;

                this._logger.LogWarning("Running search");
                this.SearchInProgress = true;

                NotifyStateChanged();

                var searchResult = await this._httpClient.GetJsonAsync<TeamSearchResponse>($"http://localhost:8080/team?search={searchTerm}");

                if (string.IsNullOrEmpty(searchResult.Err))
                {
                    this.TeamSearchResults = searchResult.Teams;
                }

                this.SearchInProgress = false;

                NotifyStateChanged();
            }
            catch (Exception ex)
            {
                this._logger.LogError(ex, ex.Message);
                this._logger.LogError(ex, "Failure running search");
            }
        }
        public async Task CreateTeam(Team team)
        {
            try
            {
                this._logger.LogInformation($"{team.Name}");

                var httpContent = new StringContent(JsonSerializer.Serialize(team), Encoding.UTF8, "application/json");

                var postResult = await this._httpClient.PostAsync($"http://localhost:8080/team", httpContent);

                this._logger.LogInformation(postResult.StatusCode.ToString());

                if (postResult.IsSuccessStatusCode)
                {
                    NotifyStateChanged();
                }
                else
                {
                    throw new Exception(await postResult.Content.ReadAsStringAsync());
                }
            }
            catch (Exception ex)
            {
                this._logger.LogError(ex.Message);
                this._logger.LogError("Failure adding team");
            }
        }

        private void NotifyStateChanged() => OnChange?.Invoke();
    }
}