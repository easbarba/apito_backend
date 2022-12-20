package dev.easbarba.apito_backend.controller;

import java.util.Optional;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import dev.easbarba.apito_backend.domain.Referee;
import dev.easbarba.apito_backend.service.RefereeService;

@RestController
@RequestMapping(path = "/api/v1/referees", produces = "application/json")
@CrossOrigin(origins = { "http://localhost:8080" })
class RefereeController {
  private RefereeService refereeService;

  public RefereeController(RefereeService refereeService) {
    this.refereeService = refereeService;
  }

  @PostMapping(consumes = "application/json", produces = "application/json")
  public ResponseEntity<Referee> create(@RequestBody Referee referee) {
    return refereeService.save(referee);
  }

  @GetMapping
  public ResponseEntity<Iterable<Referee>> getAll() {
    return refereeService.findAll();
  }

  @GetMapping("{id}")
  public ResponseEntity<Optional<Referee>> getOne(@PathVariable(value = "id") Long id) {
    return refereeService.findById(id);
  }

  @PutMapping("{id}")
  public ResponseEntity<?> update(@PathVariable(value = "id") Long id, @RequestBody Referee referee) {
    return refereeService.update(referee, id);
  }

  @PatchMapping("{id}")
  public ResponseEntity<?> updatePartial(@PathVariable(value = "id") Long id, @RequestBody Referee patch) {
    return refereeService.updatePartial(patch, id);
  }

  @DeleteMapping("{id}")
  public ResponseEntity<?> delete(@PathVariable(value = "id") Long id) {
    return refereeService.delete(id);
  }
}
