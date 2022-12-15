package dev.easbarba.apito_backend.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import dev.easbarba.apito_backend.domain.Referee;

@Repository
public interface RefereeRepository extends CrudRepository<Referee, Long> {

}
