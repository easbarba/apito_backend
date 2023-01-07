package dev.easbarba.apito_backend.domain;

import java.util.Arrays;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "matches")
public class Match {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    private String[] teams;
    private String score;
    private Referee referee;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public String[] getTeams() {
        return teams;
    }

    public void setTeams(String[] teams) {
        this.teams = teams;
    }

    public String getScore() {
        return score;
    }

    public void setScore(String score) {
        this.score = score;
    }

    public Referee getReferee() {
        return referee;
    }

    public void setReferee(Referee referee) {
        this.referee = referee;
    }

    public Match(long id, String[] teams, String score, Referee referee) {
        this.id = id;
        this.teams = teams;
        this.score = score;
        this.referee = referee;
    }

    @Override
    public String toString() {
        return "Match [id=" + id + ", referee=" + referee + ", score=" + score + ", teams=" + Arrays.toString(teams)
                + "]";
    }

    @Override
    public int hashCode() {
        final int prime = 31;
        int result = 1;
        result = prime * result + (int) (id ^ (id >>> 32));
        result = prime * result + ((referee == null) ? 0 : referee.hashCode());
        result = prime * result + ((score == null) ? 0 : score.hashCode());
        result = prime * result + Arrays.hashCode(teams);
        return result;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj)
            return true;
        if (obj == null)
            return false;
        if (getClass() != obj.getClass())
            return false;
        Match other = (Match) obj;
        if (id != other.id)
            return false;
        if (referee == null) {
            if (other.referee != null)
                return false;
        } else if (!referee.equals(other.referee))
            return false;
        if (score == null) {
            if (other.score != null)
                return false;
        } else if (!score.equals(other.score))
            return false;
        if (!Arrays.equals(teams, other.teams))
            return false;
        return true;
    }
}
