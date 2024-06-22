package contractor


func CreateFrequency(input *NewFrequencyInput) (*Frequency, error) {
    freq, err := NewFrequency(input)
    if err != nil {
        log.Err(err).Msg("Error creating frequency")
        return nil, err
    }

    db.Create(&freq)

    log.Trace().Interface("frequency", freq).Msg("Created frequency")
    log.Info().Msg("Created frequency")

    return freq, nil
}
