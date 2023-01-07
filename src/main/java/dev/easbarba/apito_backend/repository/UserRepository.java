package dev.easbarba.apito_backend.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import dev.easbarba.apito_backend.domain.User;

@Repository
public interface UserRepository extends CrudRepository<User, Long> {

}
